package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a prometheus counter to keep track of ping requests.
	numPings := promauto.NewCounter(prometheus.CounterOpts{
		Name: "pingapp_pings_total",
		Help: "The total number of incoming ping requests",
	})

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		numPings.Inc()
		w.Write([]byte("pong!\n"))
	}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
