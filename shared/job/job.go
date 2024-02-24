package job

import "time"

type JobTask func() error

type JobWorkStatus struct {
	Id      int
	Type    string
	Status  string
	Message string
}

type Job struct {
	Id    int
	Start time.Time
	End   time.Time
	Task  JobTask
}

func NewJob(id int, JobTask JobTask) *Job {
	return &Job{
		Id:   id,
		Task: JobTask,
	}
}
