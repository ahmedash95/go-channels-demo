package main

import (
	"github.com/ahmedash95/go-channels/metrics"
	"github.com/ahmedash95/go-channels/queue"
)

func main() {
	queue.InitQueueDispatcher()
	server := InitWebServer()
	queue.InitMetrics()

	metrics.InitPrometheus()
	server.Run()
}
