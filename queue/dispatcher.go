package queue

//JobQueue ... a buffered channel that we can send work requests on.
var JobQueue chan Queuable

//Queuable ... interface of Queuable Job
type Queuable interface {
	Handle() error
}

//Dispatcher ... worker dispatcher
type Dispatcher struct {
	maxWorkers int
	WorkerPool chan chan Queuable
	Workers    []Worker
}

//NewDispatcher ... creates new queue dispatcher
func NewDispatcher(maxWorkers int) *Dispatcher {
	// make job
	if JobQueue == nil {
		JobQueue = make(chan Queuable, 10)
	}
	pool := make(chan chan Queuable, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

//Run ... starts work of dispatcher and creates the workers
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		RunningWorkers.WithLabelValues("Emails").Inc()
		worker := NewWorker(d.WorkerPool)
		worker.Start()
		d.Workers = append(d.Workers, worker)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			RunningJobs.WithLabelValues("Emails").Inc()
			go func(job Queuable) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
