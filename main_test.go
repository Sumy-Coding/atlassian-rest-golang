package main

import (
	"fmt"
	"testing"
)

func myFunc() {
	fmt.Println("Inside my goroutine")
}

func Test2(t *testing.T) {
	//
	fmt.Println("Hello World")
	go myFunc()
	fmt.Println("Finished Execution")
}
