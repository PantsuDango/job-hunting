package result

import "job-hunting/model/params"

type ListJob struct {
	Count   int                `json:"Count"`
	ListJob []params.JobParams `json:"ListJob"`
}
