package main

import (
	"fmt"
	"time"
)

func main() {
	//创建创世区块
	genesisBlock := block{"0000000000000000000000000000000000000000000000000000000000000000", "", time.Now().Format("2006-01-02 15:04:05"), "我是创世区块", 1, "0000000000"}
	genesisBlock.hash = getHash(&genesisBlock)
	//把创世区块添加进区块链
	blockchain = append(blockchain, genesisBlock)
	fmt.Println(blockchain[0])
	i := 0
	for {
		time.Sleep(time.Second)
		newBlock := generateNewBlock(&blockchain[i], "我是区块内容", "00000")
		blockchain = append(blockchain, *newBlock)
		fmt.Println(blockchain[i+1])
		i++
	}
}
