package arrayHelper

import (
	"log"
	"testing"
)

// InArray 检查切片中是否包含指定的元素
func TestInArray(t *testing.T) {

	log.Println("检查切片中是否包含指定的元素")
	slice := []int{1, 2, 3, 4, 5}
	element := 3
	if InArray(slice, element) {
		log.Println("Element found")
	} else {
		log.Println("Element not found")
	}
}

// InArrayWithBinarySearch 检查切片中是否包含指定的元素，使用二分查找
func TestInArrayWithBinarySearch(t *testing.T) {

	log.Println("检查切片中是否包含指定的元素，使用二分查找")
	slice := []int{1, 2, 3, 4, 5}
	element := 3
	if InArrayWithBinarySearch(slice, element) {
		log.Println("Element found")
	} else {
		log.Println("Element not found")
	}
}
