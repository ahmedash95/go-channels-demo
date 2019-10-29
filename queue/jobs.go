package queue

import (
	"fmt"
	"log"
	"time"

	"github.com/ahmedash95/go-channels/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	JobsProcessed  *prometheus.CounterVec
	RunningJobs    *prometheus.GaugeVec
	ProcessingTime *prometheus.HistogramVec
)

//JobQueue ... a buffered channel that we can send work requests on.
var JobQueue chan Queuable

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

func InitMetrics() {
	JobsProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "worker",
			Subsystem: "jobs",
			Name:      "processed_total",
			Help:      "Total number of jobs processed by the workers",
		},
		[]string{"worker_id", "type"},
	)

	RunningJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "worker",
			Subsystem: "jobs",
			Name:      "running",
			Help:      "Number of jobs inflight",
		},
		[]string{"type"},
	)

	ProcessingTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "worker",
			Subsystem: "jobs",
			Name:      "process_time_seconds",
			Help:      "Amount of time spent processing jobs",
		},
		[]string{"worker_id", "type"},
	)

	metrics.PushRegister(ProcessingTime, RunningJobs, JobsProcessed)
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
	go func() {
		w.quit <- true
	}()
}
