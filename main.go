package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) >= 2 {
		plate := args[1]
		url := request(&plate)
		callbackPlayMpv(url)
	} else {
		fmt.Print(help())
	}
}
