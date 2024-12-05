package arrayHelper

import (
	"fmt"
	"reflect"
	"sync"
)

// DeduplicateSlice 移除一维 slice 中的重复元素
func DeduplicateSlice[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	seen := make(map[T]struct{}, len(slice))

	for _, elem := range slice {
		if _, exists := seen[elem]; !exists {
			seen[elem] = struct{}{}
			result = append(result, elem)
		}
	}
	return result
}

// DeduplicateNDimSlice 移除 N 维泛型 slice/数组中的重复元素，线程安全
func DeduplicateNDimSlice[T any](slice []T) []T {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		panic("输入必须是 slice 或 array 类型")
	}

	seen := sync.Map{}
	ch := make(chan reflect.Value, val.Len())
	var wg sync.WaitGroup

	wg.Add(val.Len())
	for i := 0; i < val.Len(); i++ {
		go func(index int) {
			defer wg.Done()
			elem := val.Index(index)
			key := generateKey(elem.Interface())
			if _, loaded := seen.LoadOrStore(key, struct{}{}); !loaded {
				ch <- elem
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := reflect.MakeSlice(reflect.TypeOf(slice), 0, val.Len())
	for elem := range ch {
		result = reflect.Append(result, elem)
	}

	return result.Interface().([]T)
}

// RemoveFromArray 兼容老函数名称
func RemoveFromArray[T comparable](slice []T, element T) []T {
	return RemoveElementSlice(slice, element)
}

// RemoveElementSlice 从一维 slice 中移除指定元素，线程安全
func RemoveElementSlice[T comparable](slice []T, element T) []T {
	result := make([]T, 0, len(slice))
	for _, elem := range slice {
		if elem != element {
			result = append(result, elem)
		}
	}
	return result
}

// RemoveElementNDimSlice 从 N 维 slice 中移除指定元素，线程安全
func RemoveElementNDimSlice(slice interface{}, element interface{}) interface{} {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic("输入必须是 slice 类型")
	}

	ch := make(chan reflect.Value, val.Len())
	var wg sync.WaitGroup

	wg.Add(val.Len())
	for i := 0; i < val.Len(); i++ {
		go func(index int) {
			defer wg.Done()
			elem := val.Index(index)
			if isNestedSlice(elem) {
				nestedResult := RemoveElementNDimSlice(elem.Interface(), element)
				ch <- reflect.ValueOf(nestedResult)
			} else {
				if !reflect.DeepEqual(elem.Interface(), element) {
					ch <- elem
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := reflect.MakeSlice(val.Type(), 0, val.Len())
	for elem := range ch {
		result = reflect.Append(result, elem)
	}

	return result.Interface()
}

// generateKey 为任意类型生成唯一键，保证线程安全的去重逻辑
func generateKey(value interface{}) string {
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var keys []string
		for i := 0; i < val.Len(); i++ {
			keys = append(keys, generateKey(val.Index(i).Interface()))
		}
		return fmt.Sprintf("%v", keys)
	default:
		return fmt.Sprintf("%v", value)
	}
}

// isNestedSlice 判断值是否是嵌套的 slice
func isNestedSlice(val reflect.Value) bool {
	return val.Kind() == reflect.Slice
}
