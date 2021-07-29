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

func readCommandLine() {
	// indexing of args starts at 0.
	// the 0th index contains the name of the executable.
	// if we directly run the app, go will create an exe implicitly and calls that with the fun exe name,
	// so you will end up having the full name printed
	//eg: [C:\Users\VAIBHA~1.BAN\AppData\Local\Temp\go-build3814485168\b001\exe\gsl.exe]
	args := os.Args

	fmt.Println(args)
}

func main() {
	//readInput()
	//readCommandLine()
	calculateTotalMealAmount()

}
