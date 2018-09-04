package main

import (
	"fmt"
	"github.com/woat/supatc/atc"
	"github.com/woat/supatc/inv"
)

func main() {
	fmt.Println("Main running.")
	l := inv.Retrieve(inv.StdDl())
	// 3 because I don't want to get banned.
	atc.Execute(l[:3])
}
