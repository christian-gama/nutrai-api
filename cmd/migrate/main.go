package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	cmd.Execute()
	elapsed := time.Since(start)

	fmt.Printf("Finished in %s\n", elapsed)
}
