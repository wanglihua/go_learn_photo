package learn_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Bird struct {
	Name           string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying...")
}

func TestStructFields1(t *testing.T) {
	sparrow := Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow)
	type_of_t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, type_of_t.Field(i).Name, f.Type(), f.Interface())
	}
}

func TestStructFields2(t *testing.T) {
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	type_of_t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, type_of_t.Field(i).Name, f.Type(), f.Interface())
	}
}
