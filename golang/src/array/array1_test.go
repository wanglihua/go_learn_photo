package learn_array

import (
	"fmt"
	"testing"
)

func TestArray1(t *testing.T) {
	var a = [3]int{1, 2, 3}
	var b = a
	b[1]++
	fmt.Println(a, b) //[1 2 3] [1 3 3]
}

func TestArray2(t *testing.T) {
	var a = [3]int{1, 2, 3}
	var b = &a
	b[1]++
	fmt.Println(a, b) //[1 3 3] &[1 3 3]
}
