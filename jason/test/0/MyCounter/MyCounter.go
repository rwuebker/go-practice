package MyCounter

import (
	"fmt"
)

func Mycounter() func() {
	theCount := 0
	increment := func() {
		theCount++
		fmt.Println("The count is", theCount)
	}
	return increment
}
