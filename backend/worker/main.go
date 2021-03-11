package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"iriz/internal/metrics"
	"iriz/internal/worker"

	"github.com/alitto/pond"
)

func main() {
	pool := pond.New(
		100,
		1000,
		pond.MinWorkers(10),
		pond.IdleTimeout(10*time.Second),
		pond.Strategy(pond.Balanced()),
	)
	defer pool.StopAndWait()

	workerName := flag.String("worker", "", "Worker name to start")
	flag.Parse()

	worker := worker.GetWorker(*workerName)

	go func(pool *pond.WorkerPool) {
		for {
			pool.Submit(func() {
				fmt.Println(worker.GetName())
				time.Sleep(500 * time.Millisecond)
			})
		}
	}(pool)

	http.Handle("/metrics", metrics.WorkerMetricsHandler(pool))
	http.ListenAndServe(":8080", nil)
}
