package queue

//Queuable ... interface of Queuable Job
type Queuable interface {
	Handle() error
}

//Dispatcher ... worker dispatcher
type Dispatcher struct {
	maxWorkers int
	WorkerPool chan chan Queuable
}

//NewDispatcher ... creates new queue dispatcher
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Queuable, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

//Run ... starts work of dispatcher and creates the workers
func (d *Dispatcher) Run() {
	// make job
	JobQueue = make(chan Queuable, 10)
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
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
