package job

import (
	"time"
)

type WorkerCmd struct {
	Id      int
	Command string
}

type Worker struct {
	Id       int
	JobCh    chan *Job
	CmdCh    chan *WorkerCmd
	StatusCh chan *JobWorkStatus
}

func NewWorker(id int, job chan *Job, cmd chan *WorkerCmd, status chan *JobWorkStatus) *Worker {

	return &Worker{
		Id:       id,
		JobCh:    job,
		CmdCh:    cmd,
		StatusCh: status,
	}
}

func (o *Worker) Start() {

	go func() {
		for {
			select {
			case job := <-o.JobCh:

				o.StatusCh <- &JobWorkStatus{Id: job.Id, Type: JW_TYPE_JOB, Status: JW_STATUS_START}
				job.Start = time.Now()
				err := job.Task()
				if err != nil {
					o.StatusCh <- &JobWorkStatus{Id: job.Id, Type: JW_TYPE_JOB, Status: JW_STATUS_ERROR, Message: err.Error()}
				}
				job.End = time.Now()
				o.StatusCh <- &JobWorkStatus{Id: job.Id, Type: JW_TYPE_JOB, Status: JW_STATUS_END}

			case cmd := <-o.CmdCh:

				if cmd.Id == o.Id || o.Id == 0 {
					if cmd.Command == JW_CMD_QUIT {
						o.StatusCh <- &JobWorkStatus{Id: o.Id, Type: JW_TYPE_WORKER, Status: JW_STATUS_QUIT}
						return
					}
				}
			}
		}
	}()
}
