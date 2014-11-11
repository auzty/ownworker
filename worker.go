package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func NewWorker(id int, workerQueue chan chan WorkRequest, c redis.Conn) Worker {

	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
		conn:        make(chan redis.Conn)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
	conn        chan redis.Conn
}

func (w Worker) Start() {
	go func() {
		for {

			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				fmt.Printf("worker%d: Received Work request, delaying for %f seconds\n", w.ID, work.Delay)

				time.Sleep(10)

				fmt.Printf("worker%d: Helooo, %d!\n", w.ID, work.Jumlah)

			case <-w.QuitChan:
				fmt.Printf("worker%d stoppping\n", w.ID)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
