package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	Timestamp     int64
	Nonce         int
}

//func (b *Block) setHash() {
//	timeStamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	headers := bytes.Join([][]byte{b.Data, b.PrevBlockHash, timeStamp}, []byte{})
//	hash := sha256.Sum256(headers)
//	b.Hash = hash[:]
//}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Timestamp:     time.Now().Unix(),
		Hash:          []byte{},
		Nonce:         0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	_ = encoder.Encode(b)
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	_ = decoder.Decode(&block)
	return &block
}
