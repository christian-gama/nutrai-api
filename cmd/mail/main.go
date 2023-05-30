package main

import (
	"fmt"
)

func init() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	cmd.Execute()
}
