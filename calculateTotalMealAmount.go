package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculateTotalMealAmount() {

	args := os.Args
	args = args[1:]

	if len(args) != 2 {
		if len(args) == 1 && args[0] == "/help" {
			fmt.Println("enter two argumnets, <meal amount> <tip amount in percentage>")
		} else {
			fmt.Println("use /help to get the info ")
		}

	} else {
		meal, err1 := strconv.ParseFloat(args[0], 32)
		tip, err2 := strconv.ParseFloat(args[1], 32)
		if err1 != nil || err2 != nil {
			fmt.Println("enter valid value for argumnets")
			return
		}
		fmt.Printf("\ntotal meal amount is: %2f \n", meal+(meal*tip*.01))
	}

}
