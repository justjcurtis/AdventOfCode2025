package utils

import "runtime"

func CreateWorkerPool[T any, U any](totalJobs int, worker func(<-chan T, chan<- U)) (chan<- T, <-chan U) {
	numWorkers := runtime.NumCPU() - 1
	results := make(chan U, totalJobs)
	jobs := make(chan T, totalJobs)
	for i := 0; i < numWorkers; i++ {
		go worker(jobs, results)
	}
	return jobs, results
}
