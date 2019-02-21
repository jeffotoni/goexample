// Go Api server
// @jeffotoni
// 2019-01-04

package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

    inFlightGauge := prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "myapiserver_http_in_flight_requests",
        Help: "A gauge of requests currently being served by the wrapped handler.",
    })

    counter := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "myapiserver_http_api_requests_total",
            Help: "A counter for requests to the wrapped handler.",
        },
        []string{"code", "method"},
    )

    // duration is partitioned by the HTTP method and handler. It uses custom
    // buckets based on the expected request duration.
    duration := prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "myapiserver_http_request_duration_seconds",
            Help:    "A histogram of latencies for requests.",
            Buckets: []float64{0.005, 0.01, 0.025, 0.05, .1, .25, .5, 1, 2.5, 5, 10},
        },
        []string{"code", "handler", "method"},
    )

    // responseSize has no labels, making it a zero-dimensional
    // ObserverVec.
    responseSize := prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "myapiserver_http_response_size_bytes",
            Help:    "A histogram of response sizes for requests.",
            Buckets: []float64{1, 10, 50, 100, 200, 500, 900, 1500},
        },
        []string{"code", "handler", "method"},
    )

    // Create the handlers that will be wrapped by the middleware.
    pushHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("push jeff"))
    })
    pullHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("X-Server", "ApiServer")
        w.Header().Add("Content-Type", "application/json")
        w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
        w.Header().Add("Pragma", "no-cache")
        w.Header().Add("Expires", "0")
        w.Header().Add("Authorization", "39393x38383c8383883dd939393939ddd9393")
        w.WriteHeader(http.StatusOK)
        //io.WriteString(w, "Pull")
        w.Write([]byte("Pullx"))
    })

    // Register all of the metrics in the standard registry.
    prometheus.MustRegister(inFlightGauge, counter, duration, responseSize)

    // Instrument the handlers with all the metrics, injecting the "handler"
    // label by currying.
    pushChain := promhttp.InstrumentHandlerInFlight(inFlightGauge,
        promhttp.InstrumentHandlerDuration(duration.MustCurryWith(prometheus.Labels{"handler": "push"}),
            promhttp.InstrumentHandlerCounter(counter,
                promhttp.InstrumentHandlerResponseSize(responseSize.MustCurryWith(prometheus.Labels{"handler": "push"}), pushHandler),
            ),
        ),
    )
    pullChain := promhttp.InstrumentHandlerInFlight(inFlightGauge,
        promhttp.InstrumentHandlerDuration(duration.MustCurryWith(prometheus.Labels{"handler": "pull"}),
            promhttp.InstrumentHandlerCounter(counter,
                promhttp.InstrumentHandlerResponseSize(responseSize.MustCurryWith(prometheus.Labels{"handler": "pull"}), pullHandler),
            ),
        ),
    )

    //http.Handle("/metrics", promhttp.Handler())
    http.Handle("/push", pushChain)
    http.Handle("/pull", pullChain)

    go func() {
        log.Println("Run Apiserver Addr:8089")
        if err := http.ListenAndServe(":8089", nil); err != nil {

            log.Fatal(err)
        }
    }()

    go func() {
        log.Println("metrics listening at :5022")
        if err := http.ListenAndServe(":5022", promhttp.Handler()); err != nil {
            log.Fatal(err)
        }
    }()

    // Wait until some signal is captured.
    sigC := make(chan os.Signal, 1)
    signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
    <-sigC
}
