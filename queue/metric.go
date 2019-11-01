package queue

import (
	"github.com/ahmedash95/go-channels/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	JobsProcessed  *prometheus.CounterVec
	RunningJobs    *prometheus.GaugeVec
	ProcessingTime *prometheus.HistogramVec
	RunningWorkers *prometheus.GaugeVec
)

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

	RunningWorkers = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "worker",
			Subsystem: "workers",
			Name:      "running",
			Help:      "Number of workers inflight",
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

	metrics.PushRegister(ProcessingTime, RunningJobs, JobsProcessed, RunningWorkers)
}
