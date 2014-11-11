package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var WorkQueue = make(chan WorkRequest, 100)

func Collector(c redis.Conn, key string) {

	jumlahque, err := redis.Int(c.Do("ZCOUNT", key, "-inf", "+inf"))

	if err != nil {
		fmt.Println("errorrr", err)
	}
	delay, err := time.ParseDuration("3s")
	work := WorkRequest{Jumlah: jumlahque, Delay: delay}
	WorkQueue <- work
	fmt.Println("Work request queued")
	return
}
