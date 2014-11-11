package main

import (
	"fmt"
	"time"
)

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {

	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

func (w Worker) Start() {
	go func() {
		for {

			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				fmt.Printf("worker%d: Received Work request, delaying for %f seconds\n", w.ID, work.Delay)

				time.Sleep(work.Delay)
				fmt.Printf("worker%d: Helooo, %s!\n", w.ID, work.Name)

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
