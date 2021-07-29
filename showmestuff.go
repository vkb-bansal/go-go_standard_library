package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func readInput() {
	fmt.Printf("hello we are running go: %v\n", runtime.Version())
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("whats ur name")
	//provide \n to indicate the delimeter, till \n the reader going to use
	//IMP: text will also contain the last \n, so when we just print it, we will have a line at the end of the text.
	text, _ := reader.ReadString('\n')
	fmt.Print(text)
	//display the host os
	fmt.Println("the host os is", runtime.GOOS)
}

func main() {
	readInput()
}
