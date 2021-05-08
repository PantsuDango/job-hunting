package result

import "job-hunting/model/params"

type ListJob struct {
	ListJob []params.JobParams `json:"ListJob"`
	Count   int                `json:"Count"`
}
