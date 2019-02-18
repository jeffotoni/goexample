package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	prommiddleware "github.com/slok/go-prometheus-middleware"
)

func Middleware(endpoint string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//job.Run()
		//jobsDurationHistogram.WithLabelValues(job.Type()).Observe(duration.Seconds())
		//jobsDurationHistogram.WithLabelValues(fmt.Sprintf("%d", http.StatusOK)).Observe(float64(time.Since(start).Seconds()))
		//httpRequestsResponseTime.Observe(float64(time.Since(start).Seconds()))

		log.Printf(
			"\033[5m%s \033[0;103m%s\033[0m \033[0;93m%s\033[0m \033[1;44m%s\033[0m",
			r.Method,
			r.RequestURI,
			endpoint,
			time.Since(start),
		)

		next.ServeHTTP(w, r)
	})
}

func main() {

	// Create our middleware.
	mdlw := prommiddleware.NewDefault()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	mux := http.NewServeMux()

	mux.HandleFunc("/api/document", func(rw http.ResponseWriter, req *http.Request) {
		delay := int32(r.Float64() * 3000)
		time.Sleep(time.Millisecond * time.Duration(delay))
		fmt.Printf("Responed with 'Hello, prometheus document!' in %dms\n", delay)
		rw.Write([]byte("document api."))
	})

	mux.HandleFunc("/api/ocr", func(rw http.ResponseWriter, req *http.Request) {
		delay := int32(r.Float64() * 5000)
		time.Sleep(time.Millisecond * time.Duration(delay))
		fmt.Printf("pong' in %dms\n", delay)
		rw.Write([]byte("Ocr!"))
	})

	//withMetrics := mw.MetricsPrometheus(handler)
	middpromet := mdlw.Handler("", mux)

	server :=
		&http.Server{
			Addr:    ":5025",
			Handler: middpromet,
		}

	//mux.Handle("/metrics", promhttp.Handler())

	go func() {
		if err := http.ListenAndServe(":5030", promhttp.Handler()); err != nil {
			log.Printf("Eror while serving metrics: %s", err)
		}
	}()

	log.Println("Run api with prometheus in: 5025")
	log.Fatal(server.ListenAndServe())

}
