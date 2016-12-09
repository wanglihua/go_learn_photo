package learn_slice

import (
	"fmt"
	"testing"
)

func TestSlice1_1(t *testing.T) {
	old_slice := []int{1, 2, 3, 4, 5}
	new_slice := old_slice[:3]

	old_slice[1] = 10
	fmt.Println(new_slice[1]) //输出10
}

func TestSlice1_2(t *testing.T) {
	old_array := [5]int{1, 2, 3, 4, 5}
	new_slice := old_array[:3]

	old_array[1] = 10
	fmt.Println(new_slice[1]) //输出10
}
