package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers = flag.Int("n", 4, "The Number of worker to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "addrest to listen for http request")
)

func main() {
	flag.Parse()

	fmt.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers)

	fmt.Println("Registering the collector")
	http.HandleFunc("/work", Collector)

	fmt.Println("HTTP server Listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
