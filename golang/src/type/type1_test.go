package learn_type

import (
	"fmt"
	"testing"
)

func TestType1(t *testing.T) {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234

	my_print(v1, v2, v3, v4)
}

func my_print(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is an string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknow type.")
		}
	}
}

func TestType2(t *testing.T) {
	var x interface{}

	x = "Hello"

	var a = x.(string)
	fmt.Println(a)

	var b = x.(int) // panic
	fmt.Println(b)
}
