package queue

import (
	"fmt"
	"log"
	"time"
)

// counter increases every time we create a worker
var i = 0

//Worker ... simple worker that handles queueable tasks
type Worker struct {
	Name       string
	WorkerPool chan chan Queuable
	JobChannel chan Queuable
	quit       chan bool
}

//NewWorker ... creates a new worker
func NewWorker(workerPool chan chan Queuable) Worker {
	i++
	return Worker{
		Name:       fmt.Sprintf("Worker%d", i),
		WorkerPool: workerPool,
		JobChannel: make(chan Queuable),
		quit:       make(chan bool)}
}

//Start ... initiate worker to start lisntening for upcomings queueable jobs
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				startTime := time.Now()
				// we have received a work request.
				// track the total number of jobs processed by the worker
				JobsProcessed.WithLabelValues(w.Name, "Emails").Inc()
				if err := job.Handle(); err != nil {
					log.Fatal("Error in job: %s", err.Error())
				}
				RunningJobs.WithLabelValues("Emails").Dec()
				ProcessingTime.WithLabelValues(w.Name, "Emails").Observe(time.Now().Sub(startTime).Seconds())

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	RunningWorkers.WithLabelValues("Emails").Dec()
	go func() {
		w.quit <- true
	}()
}
