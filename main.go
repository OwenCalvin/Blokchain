package main

import (
	bc "blockchain"
	"fmt"
)

func main() {
	block1 := bc.Block{
		Hash: "d",
		Prev: "d",
		Data: "s",
	}
	bc := bc.Blockchain{
		Blocks: []bc.Block{block1},
	}
	fmt.Printf("Test %s", bc.Blocks[0].Hash)
}
