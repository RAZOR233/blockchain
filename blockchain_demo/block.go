// package main

// import (
// 	"crypto/sha256"
// 	"time"
// )

// type Block struct {
// 	Time     int64
// 	Data     []byte
// 	PrevHash []byte
// 	Hash     []byte
// }

// func NewBlock(data string, prevHash []byte) *Block {
// 	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}}
// 	block.SetHash()
// 	return block
// }

// func (b *Block) SetHash() {
// 	//为Block生成hash，使用sha256.Sum256(data []byte)函数
// 	hasher := sha256.New()
// 	newTime := time.Unix(b.Time, 0).Format("2006-01-02 15:04:05")
// 	hasher.Write(append(append(b.PrevHash, []byte(newTime)...), b.Data...))
// 	out1 := hasher.Sum(nil)
// 	b.Hash = out1
// }

// ex3
package main

import (
	"time"
)

type Block struct {
	Time     int64
	Data     []byte
	PrevHash []byte
	Hash     []byte
	nouce    int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	//为Block生成hash，使用sha256.Sum256(data []byte)函数
	pow := NewProofOfWork(b)
	b.nouce, b.Hash = pow.Run()
}
