package main

import (
	"fmt"
	"time"
)

// Play ground for learning golang
// need to delete later

func workers(jobs <-chan string, i int, result chan<- string) {
	for j := range jobs {
		fmt.Printf("Worker %d started Job %s\n", i, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished Job %s\n", i, j)
		result <- "Job Finished " + fmt.Sprint(i)
	}
}

func main3() {
	jobs := make(chan string, 2)
	results := make(chan string, 2)
	for i := 1; i < 3; i++ {
		go workers(jobs, i, results)
	}

	for i := 1; i < 3; i++ {
		jobs <- "Job - " + fmt.Sprint(i)
	}
	close(jobs)
	for i := 1; i < 3; i++ {
		// jobs <- "Job - " + string(i)
		fmt.Println(<-results)
	}

}
