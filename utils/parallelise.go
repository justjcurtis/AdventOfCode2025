package utils

import (
	"runtime"
	"sync"
)

func Parallelise[T any](acc func(T, T) T, fn func(int) T, maxLength int) T {
	var results T
	workerCount := min(maxLength, runtime.NumCPU()-1)
	ch := make(chan T, workerCount)
	wg := sync.WaitGroup{}
	for i := 0; i < workerCount; i++ {
		start := maxLength / workerCount * i
		end := maxLength / workerCount * (i + 1)
		if i == workerCount-1 {
			end = maxLength
		}
		if workerCount == maxLength {
			start = i
			end = i + 1
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var result T
			for j := start; j < end; j++ {
				if j == start {
					result = fn(j)
					continue
				}
				result = acc(result, fn(j))
			}
			ch <- result
		}(i)
	}
	wg.Wait()
	for i := 0; i < workerCount; i++ {
		if i == 0 {
			results = <-ch
			continue
		}
		results = acc(results, <-ch)
	}
	return results
}

func ParalleliseVoid(fn func(int), maxLength int) {
	workerCount := min(maxLength, runtime.NumCPU()-1)
	wg := sync.WaitGroup{}
	for i := 0; i < workerCount; i++ {
		start := maxLength / workerCount * i
		end := maxLength / workerCount * (i + 1)
		if i == workerCount-1 {
			end = maxLength
		}
		if workerCount == maxLength {
			start = i
			end = i + 1
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				fn(j)
			}
		}(i)
	}
	wg.Wait()
}

func ParalleliseMap[T any](fn func(int) T, maxLength int) []T {
	results := make([]T, maxLength)
	workerCount := min(maxLength, runtime.NumCPU()-1)
	wg := sync.WaitGroup{}
	for i := 0; i < workerCount; i++ {
		start := maxLength / workerCount * i
		end := maxLength / workerCount * (i + 1)
		if i == workerCount-1 {
			end = maxLength
		}
		if workerCount == maxLength {
			start = i
			end = i + 1
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				results[j] = fn(j)
			}
		}(i)
	}
	wg.Wait()
	return results
}
