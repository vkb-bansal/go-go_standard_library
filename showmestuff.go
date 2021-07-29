package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("hello we are running go: %v", runtime.Version())
}
