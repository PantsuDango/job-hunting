package controller

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"job-hunting/db"
	"job-hunting/model/params"
	"job-hunting/model/tables"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ReturnSuccess               = 0    // 请求正常
	TokenExpire                 = 5001 // Token过期
	AccessDeny                  = 5002 // 访问拒绝
	IllegalRequestParameter     = 5003 // 请求参数非法
	AccessDBError               = 5004 // 请求数据库错误
	PasswordError               = 5005 // 密码错误
	TokenEmpty                  = 5006 // 无Token
	TokenError                  = 5007 // Token 错误
	IllegalAction               = 5008 // 非法的动作请求
	IncompleteParameters        = 5009 // 参数不全
	InternalError               = 5010 // 内部返回值错误
	RetrieveServiceCallFailed   = 5011 // 检索调用失败
	MisRetrieveServiceCFG       = 5012 // 缺少检测服务URL配置
	AccessDFSFail               = 5013 // 访问DFS服务失败
	AccessRetrieveServiceFailed = 5014 // 访问检索服务失败
	DownLoadPictureFail         = 5015 // 下载图片文件失败
	PictureExist                = 5016 // 图片已经存在
	LibIDError                  = 5017 // Lib Id 错误
	BusinessIDError             = 5018 // Business Id 错误
	DeleteDataError             = 5019 // Delete data not exist
	AlbumIDError                = 5020 // Album Id does not exist.
)

type Controller struct {
	SocialDB db.SocialDB
}

type Handler func(*gin.Context, tables.User)

var HandlerMap map[string]Handler

// Module and Action map
func init() {
	Controller := new(Controller)
	HandlerMap = map[string]Handler{
		"Controller.AddJob":               Controller.AddJob,
		"Controller.ListJob":              Controller.ListJob,
		"Controller.JobInfo":              Controller.JobInfo,
		"Controller.DeliverJob":           Controller.DeliverJob,
		"Controller.UserInfo":             Controller.UserInfo,
		"Controller.ModifyUser":           Controller.ModifyUser,
		"Controller.ModifyResume":         Controller.ModifyResume,
		"Controller.ModifyEducation":      Controller.ModifyEducation,
		"Controller.ModifyJobExpectation": Controller.ModifyJobExpectation,
		"Controller.ModifyPassword":       Controller.ModifyPassword,
	}
}

// 请求成功的返回
func JSONSuccess(ctx *gin.Context, status int, r interface{}) {
	// Create new token
	var token string
	auth := ctx.DefaultQuery("Bearer", "")
	if auth != "" {
		token, _ = createToken(auth)
	}

	ctx.JSON(status, gin.H{
		"RetCode": ReturnSuccess,
		"RetMsg":  "Success",
		"Bearer":  token,
		"Response": gin.H{
			"RequestId": uuid.NewV4(),
			"Result":    r,
			"Status":    "Success",
		},
	})
}

// 请求失败的返回
func JSONFail(ctx *gin.Context, status int, retCode int, retMsg string, h gin.H) {
	// Create new token
	var token string
	auth := ctx.DefaultQuery("Bearer", "")
	if auth != "" {
		token, _ = createToken(auth)
	}

	ctx.JSON(status, gin.H{
		"RetCode": retCode,
		"RetMsg":  retMsg,
		"Bearer":  token,
		"Response": gin.H{
			"RequestId": uuid.NewV4(),
			"Error":     h,
			"Status":    "Fail",
		},
	})
}

// 主入口
func (ct Controller) Handle(ctx *gin.Context) {
	// get mod and act
	index := new(params.ModActIndex)

	var bodyBytes []byte
	// Read the Body content
	if ctx.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
	}
	// Restore the io.ReadCloser to its original state
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := ctx.ShouldBindBodyWith(&index, binding.JSON); err != nil {
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		index.Action = ctx.PostForm("Action")
		index.Module = ctx.PostForm("Module")

		if index.Action == "" {
			JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Illegal request parameter", gin.H{
				"Code":    IllegalRequestParameter,
				"Message": err.Error(),
			})
			return
		}
	}

	ct_new := new(Controller)

	if index.Action == "Login" {
		ct_new.Login(ctx)
	} else if index.Action == "Register" {
		ct_new.Register(ctx)
	} else {
		modAndAct := fmt.Sprintf("%s.%s", index.Module, index.Action)
		var user tables.User

		i := strings.Index(modAndAct, "Pre")
		if i > 0 {
			user.ID = 0
			action := index.Action[3:]
			modAndAct = fmt.Sprintf("%s.%s", index.Module, action)
		} else {
			// parse token
			auth := ctx.DefaultQuery("Bearer", "")

			if auth == "" {
				JSONFail(ctx, http.StatusOK, TokenEmpty, "Need Token. Please login again.", gin.H{
					"Code":    TokenEmpty,
					"Message": "Request need token.",
				})
				return
			}
			token, err := parseToken(auth)
			if err != nil {
				JSONFail(ctx, http.StatusOK, TokenError, "Token error.", gin.H{
					"Code":    TokenError,
					"Message": err.Error(),
				})
				return
			}
			userID, _ := strconv.Atoi(token.Id)

			user, err = ct.SocialDB.QueryUserById(userID)
			if err != nil {
				JSONFail(ctx, http.StatusOK, AccessDBError, "Access user table error", gin.H{
					"Code":    AccessDBError,
					"Message": err.Error(),
				})
				return
			}

		}

		ctx.Set("user", user)
		handler, ok := HandlerMap[modAndAct]

		if ok {
			handler(ctx, user)
		} else {
			JSONFail(ctx, http.StatusOK, IllegalAction, "Illegal action.", gin.H{
				"Code":    IllegalAction,
				"Message": fmt.Sprintf("%s is not expected", modAndAct),
			})
		}
	}
}

// 注册接口
func (Controller Controller) Register(ctx *gin.Context) {

	var operator tables.User
	if err := ctx.ShouldBindBodyWith(&operator, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	if len(operator.Nick) == 0 || len(operator.UserName) == 0 || len(operator.Password) == 0 {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": "Illegal request parameter",
		})
		return
	}

	// Create MD5 password
	operator.Salt = GetRandomString(8)
	s := operator.Password + operator.Salt
	operator.Password = fmt.Sprintf("%x", md5.Sum([]byte(s)))

	err := Controller.SocialDB.CreateUser(operator)
	if err != nil {
		index := strings.Index(err.Error(), "1062")
		if index > 0 {
			JSONFail(ctx, http.StatusOK, AccessDBError, "Register fail, UserName already exists", gin.H{
				"Code":    AccessDBError,
				"Message": err.Error(),
			})
			return
		} else {
			JSONFail(ctx, http.StatusOK, AccessDBError, "Access user table error.", gin.H{
				"Code":    AccessDBError,
				"Message": err.Error(),
			})
		}
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 登录接口
func (Controller Controller) Login(ctx *gin.Context) {

	var user tables.User
	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	userInfo, err := Controller.SocialDB.GetUserInfo(user.UserName)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access user table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	// Check the password
	s := user.Password + userInfo.Salt
	md5Password := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	if md5Password == userInfo.Password {
		// Create Token
		expiresTime := time.Now().Unix() + int64(60*60*24)
		claims := jwt.StandardClaims{
			Audience:  userInfo.UserName,
			ExpiresAt: expiresTime,
			Id:        strconv.Itoa(userInfo.ID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Scoring",
			NotBefore: time.Now().Unix(),
			Subject:   "LogIn",
		}

		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Signed string used user's sale
		var jwtSecret = []byte(userInfo.Salt)
		token, err := tokenClaims.SignedString(jwtSecret)

		if err == nil {
			ctx.Set("user", userInfo)
			JSONSuccess(ctx, http.StatusOK, gin.H{
				"Bearer":    token,
				"ID":        userInfo.ID,
				"UserName":  userInfo.UserName,
				"Nick":      userInfo.Nick,
				"Sex":       userInfo.Sex,
				"HeadImage": userInfo.HeadImage,
				"Email":     userInfo.Email,
				"Phone":     userInfo.Phone,
				"Birthday":  userInfo.Birthday,
				"Degree":    userInfo.Degree,
				"Job":       userInfo.Job,
				"Address":   userInfo.Address,
			})
			// Token鉴权失败
		} else {
			JSONFail(ctx, http.StatusOK, TokenExpire, "Token expired", gin.H{
				"Token":   "Login failed",
				"Message": err.Error(),
			})
			return
		}
		// 密码错误
	} else {
		JSONFail(ctx, http.StatusOK, PasswordError, "Password Error", gin.H{
			"Token":   "Password Error",
			"Message": "Password Error",
		})
		return
	}
}

// 创建鉴权token
func createToken(auth string) (string, error) {
	oldClaims, _ := parseToken(auth)
	if oldClaims == nil {
		return "", nil
	}
	userID, _ := strconv.Atoi(oldClaims.Id)

	Controller := new(Controller)
	user, err := Controller.SocialDB.QueryUserById(userID)
	if err != nil {
		return "", err
	}

	// Create Token
	expiresTime := time.Now().Unix() + int64(60*60*24*7)
	claims := jwt.StandardClaims{
		Audience:  user.UserName,         // User name
		ExpiresAt: expiresTime,           //
		Id:        strconv.Itoa(user.ID), // User Id
		IssuedAt:  time.Now().Unix(),     // now
		Issuer:    "Controller",          // ManualSupplier
		NotBefore: time.Now().Unix(),     // current
		Subject:   "LogIn",               // purpose
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signed string used user's sale
	var jwtSecret = []byte(user.Salt)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

// 重置鉴权token
func parseToken(token string) (*jwt.StandardClaims, error) {
	// split three part then get claim party
	payload := strings.Split(token, ".")
	bytes, e := jwt.DecodeSegment(payload[1])

	if e != nil {
		println(e.Error())
	}
	content := ""
	for i := 0; i < len(bytes); i++ {
		content += string(bytes[i])
	}
	split := strings.Split(content, ",")
	id := strings.SplitAfter(split[2], ":")
	i := strings.Split(id[1], "\\u")
	i = strings.Split(i[0], "\"")

	operator_id, err := strconv.Atoi(i[1])
	// operator_i, er := strconv.ParseInt(i[1], 16, 64)

	if err != nil {
		println(err.Error())
	}

	u, _ := db.SocialDB{}.QueryUserById(int(operator_id))
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte(u.Salt), nil
		})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			claim.Id = strconv.Itoa(int(operator_id))
			return claim, nil
		}
	}
	return nil, err
}

// 创建盐
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
