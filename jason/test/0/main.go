package main

import "test/0/MyCounter"

func main() {
	var incFunc func()
	incFunc = MyCounter()
	for i := 0; i < 5; i++ {
		incFunc() //use () to execute increment
	}
}
