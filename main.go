package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	bc := NewBlockchain()
	time.Sleep(time.Second * 3)
	bc.AddBlock("Send 1 BTC to Jung")
	time.Sleep(time.Second * 3)
	bc.AddBlock("Send 2 BTC to Jung")

	for _, block := range bc.blocks {
		fmt.Printf("Prev.Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)

		pow := NewProofOfWork(block)
		fmt.Printf("POW： %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
