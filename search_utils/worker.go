package search_utils

import (
	"log"
	"time"

	common "github.com/ilius/go-dict-commons"
)

func RunWorkers(
	count int,
	workerCount int,
	timeout time.Duration,
	worker func(int, int) []*common.SearchResultLow,
) []*common.SearchResultLow {
	if workerCount < 2 {
		return worker(0, count)
	}
	if count < 2*workerCount {
		return worker(0, count)
	}

	ch := make(chan []*common.SearchResultLow, workerCount)

	sender := func(start int, end int) {
		ch <- worker(start, end)
	}

	step := count / workerCount
	start := 0
	for range workerCount - 1 {
		end := start + step
		go sender(start, end)
		start = end
	}
	go sender(start, count)

	results := []*common.SearchResultLow{}
	timeoutCh := time.NewTimer(timeout)
	for range workerCount {
		select {
		case wRes := <-ch:
			results = append(results, wRes...)
		case <-timeoutCh.C:
			log.Println("Search Timeout")
			return results
		}
	}

	return results
}
