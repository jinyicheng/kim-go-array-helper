package arrayHelper

import (
	"runtime"
	"sync"
)

// RemoveFromArray 从切片中移除指定的元素，使用并行处理
func RemoveFromArray[T comparable](slice []T, element T) []T {
	numWorkers := runtime.GOMAXPROCS(0)
	if numWorkers > len(slice) {
		numWorkers = len(slice)
	}
	if numWorkers > 8 { // 限制最大并发数量
		numWorkers = 8
	}

	chunkSize := (len(slice) + numWorkers - 1) / numWorkers
	var wg sync.WaitGroup
	resultCh := make(chan []T, numWorkers)

	pool := sync.Pool{
		New: func() interface{} {
			return make([]T, 0, chunkSize)
		},
	}

	// 启动多个 goroutine 处理子切片
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			var newSubSlice []T
			for _, item := range slice[start:end] {
				if item != element {
					newSubSlice = append(newSubSlice, item)
				}
			}
			pool.Put(newSubSlice[:0]) // 清空切片并放回池中
			resultCh <- newSubSlice
		}(start, end)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	close(resultCh)

	// 合并结果
	var result []T
	for subResult := range resultCh {
		result = append(result, subResult...)
	}

	return result
}
