package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var collectorContainer []prometheus.Collector

//InitPrometheus ... initalize prometheus
func InitPrometheus() {
	prometheus.MustRegister(collectorContainer...)
}

//PushRegister ... Push collectores to prometheus before inializing
func PushRegister(c ...prometheus.Collector) {
	collectorContainer = append(collectorContainer, c...)
}
