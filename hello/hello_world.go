package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print(os.Args)

	if len(os.Args) > 1 {
		fmt.Println(" Hello World")
	}

	os.Exit(0)
}
