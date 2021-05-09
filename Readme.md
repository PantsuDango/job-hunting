### 公共参数

##### 简要描述

- 公共参数说明

##### 请求URL
- ` http://39.108.110.77/job-hunting/api `
##### 请求方式
- POST 

##### 返回示例 

``` 
{
    "Bearer": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZG1pbiIsImV4cCI6MTYxODg1NTYyMSwianRpIjoiMSIsImlhdCI6MTYxODc2OTIyMSwiaXNzIjoiQ29udHJvbGxlciIsIm5iZiI6MTYxODc2OTIyMSwic3ViIjoiTG9nSW4ifQ.Mrbm5tEIeLWiK49dQf9l4LqVzKcYN8rsxOCpB9Jeuds",
    "Response": {
        "RequestId": "e3fcb873-5be1-488a-b190-c3f6b97aeb40",
        "Result": {},
        "Status": "Success"
    },
    "RetCode": 0,
    "RetMsg": "Success"
}
```

##### 返回参数说明 

| 参数名   | 类型        | 说明                  |
| :------- | :---------- | --------------------- |
| Bearer   | String      | 鉴权token，有效期一周 |
| RetCode  | Int         | 状态码，0表示正常     |
| RetMsg   | String      | 状态信息              |
| Response | ResponseMap | 响应信息              |

###### Response

| 参数名    | 类型      | 说明                                   |
| :-------- | :-------- | -------------------------------------- |
| RequestId | String    | 本次请求唯一ID                         |
| Status    | String    | Success 或 Fail                        |
| Result    | Interface | 接口具体返回的详细数据，有错误则不返回 |
| Error     | String    | 错误信息，没有错误则不返回             |

##### 错误码 

| 错误码 | 说明           |
| :----- | :------------- |
| 0      | 请求正常       |
| 5001   | Token过期      |
| 5002   | 访问拒绝       |
| 5003   | 请求参数非法   |
| 5004   | 请求数据库错误 |
| 5005   | 密码错误       |
| 5006   | 无Token        |
| 5007   | Token          |
| 5008   | 非法的动作请求 |
| 5009   | 参数不全       |



### 登录


##### 简要描述

- 用户登录

##### 请求URL
- ` http://39.108.110.77/job-hunting/api `
##### 请求方式
- POST 

##### 请求示例 

``` 
{
    "Module": "Controller",
    "Action": "Login",
    "UserName": "inyin",
    "Password": "123456"
}
```

##### 请求参数

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| Module   | 是   | String | 模块名 |
| Action   | 是   | String | 动作名 |
| UserName | 是   | String | 用户名 |
| Password | 是   | String | 密码   |

##### 返回示例 

``` json
{
    "Address": "",
    "Bearer": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpbnlpbiIsImV4cCI6MTYxOTg0MTk5OCwianRpIjoiMSIsImlhdCI6MTYxOTc1NTU5OCwiaXNzIjoiU2NvcmluZyIsIm5iZiI6MTYxOTc1NTU5OCwic3ViIjoiTG9nSW4ifQ.79rpcWhmg9_SA6ATHm4s3NJpD7tbDNIZEqo-9AvLGgk",
    "Birthday": "",
    "Email": "",
    "HeadImage": "",
    "ID": 1,
    "Job": "",
    "Nick": "叶叶",
    "Phone": "",
    "Sex": 0,
    "UserName": "inyin"
}
```

##### 返回参数说明 

| 参数名    | 类型   | 说明                     |
| --------- | :----- | ------------------------ |
| Bearer    | String | 鉴权token                |
| ID        | Int    | 用户操作ID               |
| Nick      | String | 用户昵称                 |
| UserName  | String | 用户名                   |
| HeadImage | String | 头像图片链接             |
| Sex       | Int    | 性别: 0-保密, 1-男, 2-女 |
| Phone     | String | 手机号                   |
| Email     | String | 邮箱                     |
| Birthday  | String | 生日                     |
| Job       | String | 职位                     |
| Address   | String | 家庭住址                 |





### 注册


##### 简要描述

- 用户注册

##### 请求URL
- ` http://39.108.110.77/job-hunting/api `
##### 请求方式
- POST 

``` 
{
    "Module": "Controller",
    "Action": "Register",
    "Nick": "叶叶",
    "UserName": "inyin",
    "Password": "123456"
}
```
##### 参数

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| Module   | 是   | String | 模块名 |
| Action   | 是   | String | 动作名 |
| Nick     | 是   | String | 昵称   |
| UserName | 是   | String | 用户名 |
| Password | 是   | String | 密码   |

##### 返回示例 

``` 
  {
    "Success"
  }
```





### 图片上传


##### 简要描述

- 图片上传，具体可访问 http://39.108.110.77/dfs 有可视化页面测试

##### 请求URL
- ` http://39.108.110.77/upload `
##### 请求方式
- POST 

##### 参数

| 参数名 | 必选 | 说明              |
| :----- | :--- | :---------------- |
| file   | 是   | 上传的文件        |
| output | 否   | 输出，可选填 json |

##### 返回示例 

``` 
{
    "url": "http://39.108.110.77/group1/default/20210423/00/50/5/a4a2be1aaebe0b6f0946654e8c4e2061.json",
    "md5": "164cb8185fcc668749d59571c1e5a345",
    "path": "/group1/default/20210423/00/50/5/a4a2be1aaebe0b6f0946654e8c4e2061.json",
    "domain": "http://39.108.110.77",
    "scene": "default",
    "size": 217,
    "mtime": 1619110234,
    "scenes": "default",
    "retmsg": "",
    "retcode": 0,
    "src": "/group1/default/20210423/00/50/5/a4a2be1aaebe0b6f0946654e8c4e2061.json"
}
```




### 新建职位


##### 简要描述

- 新建职位

##### 请求URL
- ` http://39.108.110.77/job-hunting/api `
##### 请求方式
- POST

``` 
{
    "Module": "Controller",
    "Action": "AddJob",
    "Name": "高级web开发工程师",
    "Pay": "15-20k",
    "IcoUrl": "",
    "Company": "华融科技",
    "Scale": "500-2000人",
    "Description": "团队氛围融洽，福利健全",
    "Tags": [
        "长沙",
        "5-10年",
        "本科"
    ]
}
```
##### 参数

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| Module   | 是   | String | 模块名 |
| Action   | 是   | String | 动作名 |
| Name     | 是   | String | 职位名称   |
| Pay | 是   | String | 薪资待遇 |
| IcoUrl | 否   | String | 公司图标   |
| Company | 是   | String | 公司名称   |
| Scale | 否   | String | 公司规模   |
| Description | 否   | String | 描述   |
| Tags | 否   | StringArray | 职位标签列表   |

##### 返回示例

``` 
  {
    "Success"
  }
```




### 查询职位列表


##### 简要描述

- 查询职位列表

##### 请求URL
- ` http://39.108.110.77/job-hunting/api `
##### 请求方式
- POST

``` 
{
    "Module": "Controller",
    "Action": "ListJob",
    "Offset": 0,
    "Limit": 5,
    "Keyword": "开发"
}
```
##### 参数

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| Module   | 是   | String | 模块名 |
| Action   | 是   | String | 动作名 |
| Offset   | 否   | Int | 跳过的职位数   |
| Limit    | 否   | Int | 显示的职位数, 默认10 |
| Keyword   | 否   | String | 关键字查询 |

##### 返回示例

``` 
{
    "Count": 4,
    "ListJob": [
        {
            "ID": 4,
            "Name": "python开发工程师",
            "Pay": "13-16k",
            "IcoUrl": "",
            "Company": "智宸科技",
            "Scale": "50-200人",
            "Description": "团队氛围融洽，福利健全",
            "Tags": [
                "程度",
                "1-3年",
                "本科"
            ],
            "Createtime": "2021-05-07 19:41:23",
            "Isdeliver": false
        },
        {
            "ID": 3,
            "Name": "java开发工程师",
            "Pay": "15-20k",
            "IcoUrl": "",
            "Company": "小米科技",
            "Scale": "100-200人",
            "Description": "团队氛围融洽，福利健全",
            "Tags": [
                "杭州",
                "2-3年",
                "本科"
            ],
            "Createtime": "2021-05-07 19:40:27",
            "Isdeliver": false
        },
        {
            "ID": 2,
            "Name": "初级web开发工程师",
            "Pay": "10-12k",
            "IcoUrl": "",
            "Company": "华胜智能",
            "Scale": "50-200人",
            "Description": "团队氛围融洽，福利健全",
            "Tags": [
                "北京",
                "1-2年",
                "本科"
            ],
            "Createtime": "2021-05-07 19:39:47",
            "Isdeliver": false
        },
        {
            "ID": 1,
            "Name": "高级web开发工程师",
            "Pay": "15-20k",
            "IcoUrl": "",
            "Company": "华融科技",
            "Scale": "500-2000人",
            "Description": "团队氛围融洽，福利健全",
            "Tags": [
                "长沙",
                "5-10年",
                "本科"
            ],
            "Createtime": "2021-05-07 19:32:41",
            "Isdeliver": true
        }
    ]
}
```




### 查询某职位详情


##### 简要描述

- 查询某职位详情

##### 请求URL
- ` http://39.108.110.77/job-hunting/api `
##### 请求方式
- POST

``` 
{
    "Module": "Controller",
    "Action": "JobInfo",
    "ID": 1
}
```
##### 参数

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| Module   | 是   | String | 模块名 |
| Action   | 是   | String | 动作名 |
| ID   | 是   | Int | 该职位ID   |

##### 返回示例

``` 
{
    "ID": 1,
    "Name": "高级web开发工程师",
    "Pay": "15-20k",
    "IcoUrl": "",
    "Company": "华融科技",
    "Scale": "500-2000人",
    "Description": "团队氛围融洽，福利健全",
    "Tags": [
        "长沙",
        "5-10年",
        "本科"
    ],
    "Createtime": "2021-05-07 19:32:41",
    "Isdeliver": false
}
```




### 职位投递


##### 简要描述

- 职位投递

##### 请求URL
- ` http://39.108.110.77/job-hunting/api `
##### 请求方式
- POST

``` 
{
    "Module": "Controller",
    "Action": "DeliverJob",
    "ID": 1
}
```
##### 参数

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| Module   | 是   | String | 模块名 |
| Action   | 是   | String | 动作名 |
| ID   | 是   | Int | 该职位ID   |

##### 返回示例

``` 
{
    "Success"
}
```




### 用户个人信息


##### 简要描述

- 用户个人信息

##### 请求URL
- ` http://39.108.110.77/job-hunting/api `
##### 请求方式
- POST

``` 
{
    "Module": "Controller",
    "Action": "UserInfo"
}
```
##### 参数

| 参数名   | 必选 | 类型   | 说明   |
| :------- | :--- | :----- | ------ |
| Module   | 是   | String | 模块名 |
| Action   | 是   | String | 动作名 |

##### 返回示例

``` 
{
    "DeliverCount": 1,
    "UserInfo": {
        "ID": 1,
        "Nick": "叶叶",
        "UserName": "inyin",
        "Sex": 0,
        "HeadImage": "",
        "Email": "",
        "Phone": "",
        "Birthday": "",
        "Degree": "",
        "Job": "",
        "Address": ""
    }
}
```