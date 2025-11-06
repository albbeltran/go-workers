package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("El worker %d inicio la tarea %d\n", id, job)
		time.Sleep(300 * time.Millisecond)

		result := job * 2

		results <- result

		fmt.Printf("El worker %d finalizo la tarea %d\n", id, job)
	}
}

func main() {
	const numWorkers = 4
	const numJobs = 20

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	for r := 1; r <= numJobs; r++ {
		<-results
	}
}
