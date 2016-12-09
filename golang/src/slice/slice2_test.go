package learn_slice

import (
	"fmt"
	"testing"
)

func TestSlice2_1(t *testing.T) {
	//引用，非复制，所以任何对slice1或slice的修改都会影响对方
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	data1 := data[0:2]
	data1[0] = 99
	fmt.Println(data1)
	fmt.Println(data)

	//[99 2]
	//[99 2 3 4 5 6 7 8 9 0]
}

func TestSlice2_2(t *testing.T) {
	data := make([]int, 10, 20)
	data[0] = 1
	data[1] = 2
	dataappend := make([]int, 10, 20) //len <=10 则 	result[0] = 99 会 影响源Slice
	dataappend[0] = 1
	dataappend[1] = 2
	result := append(data, dataappend...)
	result[0] = 99
	result[11] = 98
	fmt.Println("length:", len(data), ":", data)
	fmt.Println("length:", len(result), ":", result)
	fmt.Println("length:", len(dataappend), ":", dataappend)
}
