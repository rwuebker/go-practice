package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 6.54
	v := reflect.ValueOf(x)
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("kind is float64:", v.Kind() == reflect.String)
}
