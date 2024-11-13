package arrayHelper

import (
	"log"
	"testing"
)

// InArray 检查切片中是否包含指定的元素
func TestRemoveFromArray(t *testing.T) {

	log.Println("获取指定日期对应的周第一天")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	element := 5
	log.Println(RemoveFromArray(slice, element))
}
