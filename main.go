package main

import (
	"block"
	"fmt"
)

func main() {
	b1 := block.Block{
		Hash: "d",
		Prev: "d",
		Data: "s",
	}
	fmt.Printf("Test %s", b1.Hash)
}
