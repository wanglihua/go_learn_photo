package main

import (
	"fmt"
	"os"
)

//ToDo: ddd
func main() {
	var str string = "Hello, 你好"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	fmt.Println()

	for _, c := range str {
		fmt.Println(string(c))
	}

	fmt.Println(os.Getenv("HOME"))

	say_hello()

	print("ddd")
}

func say_hello() {
	fmt.Println("Hello world!")
}
