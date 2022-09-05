package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// returns a hash of the transactions in the the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.ID)
	}
	mTree := NewMerkleTree(transactions)
	return mTree.RootNode.Data
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	// todo fixed timestamp
	//ts := time.Now().Unix()
	ts := int64(1661823364)
	block := &Block{ts, transactions, prevBlockHash, []byte{}, 0}
	// 挖矿
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
