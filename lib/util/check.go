package util

import "fmt"

func CheckPanic(err error) {
	if err != nil {
		panic(fmt.Errorf("panic: %v", err))
	}
}
