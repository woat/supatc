package main

import (
	"fmt"
	"github.com/woat/supatc/inv"
)

func main() {
	fmt.Println("Main running.")
	l := inv.Retrieve(inv.StdDl())
	fmt.Printf("Retrieve running... %v", l)
	if fnd, ok := inv.Find("black", l); ok {
		fmt.Printf("Find running... %v", fnd)
	}
}
