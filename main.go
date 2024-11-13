package arrayHelper

import (
	"golang.org/x/exp/constraints"
	"runtime"
	"sort"
	"sync"
)

// InArray 检查切片中是否包含指定的元素
func InArray[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// InArrayWithBinarySearch 检查切片中是否包含指定的元素，使用二分查找
func InArrayWithBinarySearch[T constraints.Ordered](slice []T, element T) bool {
	// 先对切片进行排序
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	// 使用二分查找
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= element
	})

	return index < len(slice) && slice[index] == element
}

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
