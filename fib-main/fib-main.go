package main

import (
	"fmt"
	"github.com/lanceman/go-experiments/fib"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			converted, err := strconv.ParseUint(val, 0, 64)
			if err != nil {
				fmt.Printf("Failed to convert [%s] to uint64\n", val)
				fmt.Println(err)
				os.Exit(0)
			}
			a,b := fib.NextFib(converted)
			fmt.Printf("Entered [%d], NextFib returned [%d]: sequence[%d]\n",converted,a,b)
		}
	} else {
		fmt.Println("No arguments provided!")
	}
}
