package workermanager

import (
	"fmt"
	"net"
)

type Job struct {
	Action func(net.Conn)
	Conn   net.Conn
}

func Start(numberOfWorkers int, jobs <-chan Job) {
	for i := 1; i <= numberOfWorkers; i++ {
		go worker(i, jobs)
	}
}

func worker(i int, jobs <-chan Job) {
	for job := range jobs {
		fmt.Println("Processing Job on worker", i)
		job.Action(job.Conn)
	}
}
