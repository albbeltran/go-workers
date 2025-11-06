package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

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

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := 1; r <= numJobs; r++ {
		<-results
	}
}
