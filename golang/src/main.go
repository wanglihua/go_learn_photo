package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println()

	var once sync.Once
	once.Do(Test)
}

func Test() {
	fmt.Println("test")
}
