package arrayHelper

import (
	"golang.org/x/exp/constraints"
	"sort"
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
