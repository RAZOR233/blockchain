package main

import (
	"log"
	"time"
)

type block struct {
	//上一个区块的 Hash
	Lasthash string
	//本区块 Hash
	Hash string
	//区块存储的数据（比如比特币 UTXO 模型 则此处可用于存储交易）
	Data string
	//时间戳
	Timestamp string
	//区块高度
	Height int
	//难度值
	DiffNum uint
	//随机数
	Nonce int
}

// 区块挖矿（通过自身递增 nonce 值计算 hash）
var blockchain []block

func NewBlock(last *block, data string) *block {
	a := new(block)
	a.Data = data
	a.DiffNum = 6
	a.Height = last.Height + 1
	a.Lasthash = last.Hash
	a.Timestamp = time.Now().String()
	return a
}
func getHash(newBlock *block) *block {

	pow := NewProofOfWork(newBlock)
	newBlock.Nonce, newBlock.Hash = pow.Run()
	return newBlock
}
func mine(data string) block {
	if len(blockchain) < 1 {
		log.Panic("还未生成创世区块！")
	}
	lastBlock := blockchain[len(blockchain)-1]
	//制造一个新的区块
	newBlock := NewBlock(&lastBlock, data)
	newBlock = getHash(newBlock)
	return *newBlock
}
