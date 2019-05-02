package main

import (
	"github.com/liamg/tml"
)

func main() {

	tml.Printf(
		"<red>%s</red> <yellow>%s</yellow> <green><bold>%s</bold></green>\n",
		"ERROR",
		"WARNING",
		"SUCCESS",
	)

}
