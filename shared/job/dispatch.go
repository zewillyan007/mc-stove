package job

import (
	"context"
	"errors"
	"mc-stove/shared/port"
	"sync"
)

type Dispatcher struct {
	jobCounter int
	ctx        context.Context
	jobCh      chan *Job
	workCh     chan *Job
	workCmdCh  chan *WorkerCmd
	statusCh   chan *JobWorkStatus
	running    bool
	log        port.ILogger
	sync.RWMutex
}

func NewDispatcher(log port.ILogger, ctx context.Context) *Dispatcher {

	return &Dispatcher{
		jobCounter: 0,
		ctx:        ctx,
		jobCh:      make(chan *Job),
		workCh:     make(chan *Job),
		workCmdCh:  make(chan *WorkerCmd),
		statusCh:   make(chan *JobWorkStatus),
		running:    false,
		log:        log.SetExtraPart("fromp", "worker"),
		RWMutex:    sync.RWMutex{},
	}
}

func (o *Dispatcher) AddJob(jobt JobTask) {

	job := NewJob(o.jobCounter, jobt)
	go func() { o.jobCh <- job }()
	o.jobCounter++
}

func (o *Dispatcher) Finished() bool {

	o.RLock()
	defer o.RUnlock()

	if o.jobCounter < 1 {
		return true
	} else {
		return false
	}
}

func (o *Dispatcher) Running() bool {
	return o.running
}

func (o *Dispatcher) Start(numWorkers int) error {

	if numWorkers < 1 {
		return errors.New("Start requires >= 1 workers")
	}

	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i, o.workCh, o.workCmdCh, o.statusCh)
		worker.Start()
	}

	o.running = true

	go func() {
		for {
			select {

			case job := <-o.jobCh:

				go func() { o.workCh <- job }()

			case status := <-o.statusCh:

				if status.Status == JW_STATUS_ERROR {
					o.log.Warn("%v-%v: %v", status.Type, status.Id, status.Message)
				}
			}
		}
	}()

	return nil
}

func (o *Dispatcher) Close() {
	close(o.jobCh)
	close(o.workCh)
	close(o.statusCh)
	close(o.workCmdCh)
	o.log.Info("%v", "Close Worker")
}
