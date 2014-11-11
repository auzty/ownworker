package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	NWorkers  = flag.Int("n", 4, "The Number of worker to start")
	conn, err = redis.Dial("tcp", "localhost:6379")
)

func main() {
	flag.Parse()

	fmt.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers, conn)

	fmt.Println("Registering the collector")
	Collector(conn, "resque:gitlab:schedule")

	defer conn.Close()
}
