package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/liamg/tml"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	tml.Printf(string(input))
}
