package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

type block struct {
	//上一个块的 hash
	prehash string
	//本块 hash
	hash string
	//时间戳
	timestamp string
	//区块内容
	data string
	//区块高度
	height int
	//挖出本块的地址
	address string
}

// 代表挖矿节点
type node struct {
	//代币数量
	tokens int
	//质押时间
	days int
	//节点地址
	address string
}

var mineNodesPool []node
var probabilityNodesPool []node
var blockchain []block

func prepareData(block *block) []byte {
	lasthash, _ := hex.DecodeString(block.prehash)
	data := bytes.Join(

		[][]byte{
			lasthash,
			[]byte(block.data),
			[]byte(block.timestamp),
			[]byte(block.address),
		},
		[]byte{},
	)

	return data
}

// 初始化
func init() {
	rand.Seed(time.Now().UnixNano())
	//手动添加两个节点
	mineNodesPool = append(mineNodesPool, node{1000, 1, "AAAAAAAAAA"})
	mineNodesPool = append(mineNodesPool, node{100, 3, "BBBBBBBBBB"})
	//初始化随机节点池（挖矿概率与代币数量和币龄有关）
	for _, v := range mineNodesPool {
		for i := 0; i <= v.tokens*v.days; i++ {
			probabilityNodesPool = append(probabilityNodesPool, v)
		}
		//fmt.Println(len(probabilityNodesPool))
	}
}
func getMineNodeAddress() string {
	return probabilityNodesPool[rand.Intn(len(probabilityNodesPool))].address
}

func getHash(a *block) string {
	hasher := sha256.New()
	hasher.Write(prepareData(a))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash[:])
}
func generateNewBlock(last *block, data string, t string) *block {
	a := new(block)
	a.data = data
	a.height = last.height + 1
	a.prehash = last.hash
	a.timestamp = time.Now().String()
	a.address = getMineNodeAddress()
	a.hash = getHash(a)
	return a
}
