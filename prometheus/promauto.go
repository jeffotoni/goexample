// Go Api server
// @jeffotoni
// 2019-01-04

package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var sendCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "umovme_sms",
	Subsystem: "message_send",
	Name:      "total",
	Help:      "Count of sms messages send.",
}, []string{
	"phone",
})

func main() {
	http.Handle("/metrics", promhttp.Handler())
	sendCounter.WithLabelValues("5199999999").Inc()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
