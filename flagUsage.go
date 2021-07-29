package main

import (
	"flag"
	"fmt"
)

func flagUsage() {
	arch := flag.String("arch", "x86", "CPU type")
	flag.Parse()
	switch *arch {
	case "x86":
		fmt.Println("using x86 arc")
	case "AMD64":
		fmt.Println("using AMD64 arch")

	}
}
