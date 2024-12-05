package arrayHelper

import (
	"log"
	"testing"
)

// InArray 检查切片中是否包含指定的元素
func TestRemoveFromArray(t *testing.T) {
	log.Println("删除指定值")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	element := 5
	log.Println(RemoveFromArray(slice, element))
}

func TestRemoveElementNDimSlice(t *testing.T) {
	log.Println("删除指定值")
	// 定义一个多维 slice
	multiDimSlice := [][]int{
		{1, 2, 3, 4},
		{5, 3, 7, 8},
		{9, 10, 3, 11},
	}
	element := 5
	log.Println(RemoveElementNDimSlice(multiDimSlice, element))
}

func TestDeduplicateSlice(t *testing.T) {
	log.Println("一维 slice 去重")
	slice1D := []int{1, 2, 2, 3, 4, 4, 5}
	log.Println(DeduplicateSlice(slice1D))
}

func TestDeduplicateNDimSlice(t *testing.T) {
	log.Println("二维 slice 去重")
	// 示例：包含结构体的一维 slice 去重
	sliceStruct := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{Name: "j", Age: 25},
		{Name: "y", Age: 30},
		{Name: "c", Age: 22},
		{Name: "a", Age: 28},
		{Name: "y", Age: 30}, // 重复项
		{Name: "a", Age: 28}, // 重复项
	}
	dedupStruct := DeduplicateNDimSlice(sliceStruct)
	log.Println(DeduplicateNDimSlice(dedupStruct))
} // 定义一个结构体来表示一个人
