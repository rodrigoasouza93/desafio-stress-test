package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "", "URL to fetch")
	requests := flag.Int("requests", 1, "Number of requests to make")
	concurrency := flag.Int("concurrency", 1, "Number of concurrency to use")
	flag.Parse()

	validate(url, requests, concurrency)

	jobs := make(chan int, *requests)
	results := make(chan int, *requests)

	var wg sync.WaitGroup
	var startAt = time.Now()

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			worker(*url, jobs, results)
			wg.Done()
		}()
	}

	for i := 0; i < *requests; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	close(results)

	statusCodes := make(map[int]int)
	for result := range results {
		statusCodes[result]++
	}

	var elapsed = time.Since(startAt)
	fmt.Println("Tempo de execução:", elapsed)
	fmt.Println("Total de requisições:", *requests)
	for code, count := range statusCodes {
		if code == 0 {
			fmt.Printf("Falhas: %d\n", count)
		} else {
			fmt.Printf("Status %d: %d\n", code, count)
		}
	}
}

func validate(url *string, requests *int, concurrency *int) {
	if *url == "" {
		panic("url is required")
	}
	if *requests <= 0 {
		panic("requests must be greater than 0")
	}
	if *concurrency <= 0 {
		panic("concurrency must be greater than 0")
	}
}

func worker(url string, jobs <-chan int, results chan<- int) {
	for range jobs {
		resp, err := http.Get(url)
		if err != nil {
			results <- 0
			continue
		}
		results <- resp.StatusCode
		resp.Body.Close()
	}
}
