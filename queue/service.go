package queue

var QueueDispatcher *Dispatcher

func InitQueueDispatcher() {
	QueueDispatcher = NewDispatcher(4)
	QueueDispatcher.Run()
}

func Push(job Queuable) {
	JobQueue <- job
}
