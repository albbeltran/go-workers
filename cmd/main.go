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
