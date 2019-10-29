package queue

var QueueDispatcher *Dispatcher

func InitQueueDispatcher() {
	QueueDispatcher = NewDispatcher(2)
	QueueDispatcher.Run()
}

func Push(job Queuable) {
	JobQueue <- job
}
