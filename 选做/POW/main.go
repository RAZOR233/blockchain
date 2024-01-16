package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//制造一个创世区块
	genesisBlock := new(block)
	genesisBlock.Timestamp = time.Now().String()
	genesisBlock.Data = "我是创世区块！"
	genesisBlock.Lasthash =
		"0000000000000000000000000000000000000000000000000000000000000000"
	genesisBlock.Height = 1
	genesisBlock.Nonce = 0
	genesisBlock.DiffNum = 0
	genesisBlock = getHash(genesisBlock)
	fmt.Println(*genesisBlock)
	//将创世区块添加进区块链
	blockchain = append(blockchain, *genesisBlock)
	for i := 0; i < 10; i++ {
		newBlock := mine("天气不错" + strconv.Itoa(i))
		blockchain = append(blockchain, newBlock)
		fmt.Println(newBlock)
	}
}
