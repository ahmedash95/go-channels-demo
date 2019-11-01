package main

import (
	"github.com/ahmedash95/go-channels/metrics"
	"github.com/ahmedash95/go-channels/queue"
)

func main() {
	queue.InitMetrics()
	queue.InitQueueDispatcher()
	server := InitWebServer()

	metrics.InitPrometheus()
	server.Run()
}
