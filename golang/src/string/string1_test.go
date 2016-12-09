package learn_string

import (
	"fmt"
	"testing"
)

func TestString1(t *testing.T) {
	str := "Hello,世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}
}

func TestString2(t *testing.T) {
	str := "Hello,世界"
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}
}
