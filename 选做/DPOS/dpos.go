package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
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

// 代表竞选节点
type supernode struct {
	//得票数
	votes int
	//节点地址
	address string
}

// 代表投票节点
type votenode struct {
	//持币量
	tokens int
	//节点地址
	address string
}

var voteNodesPool []votenode
var superNodesPool []supernode
var finalNodesPool []supernode
var blockchain []block
var count int

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
	count = -1
	fmt.Println("正在初始化100个投票节点")
	for i := 0; i < 100; i++ {
		token := rand.Intn(10000)
		name := "投票节点" + strconv.Itoa(i)
		tmp := votenode{token, name}
		fmt.Print(tmp)
		fmt.Print(",")
		voteNodesPool = append(voteNodesPool, tmp)
	}
	fmt.Print("\n")
	fmt.Println("当前存在的10个竞选节点")
	for i := 0; i <= 10; i++ {
		token := 0
		name := "竞选节点" + strconv.Itoa(i)
		tmp2 := supernode{token, name}
		fmt.Print(tmp2)
		fmt.Print(",")
		superNodesPool = append(superNodesPool, tmp2)
	}
	fmt.Print("\n")
	fmt.Println("投票节点们开始进行投票...")
	vote()
	fmt.Println("投票结束，查看节点所得票数：")
	for _, v := range superNodesPool {
		fmt.Print(v)
		fmt.Print(",")
	}
	list := max()
	for i := 0; i < 3; i++ {
		finalNodesPool = append(finalNodesPool, superNodesPool[list[i]])
	}
	fmt.Println("对竞选节点按照得票数排序，前3名当选超级节点：")
	fmt.Println(finalNodesPool)
	fmt.Println("随机打乱顺序，得到")
	rand.Shuffle(len(finalNodesPool), func(i, j int) {
		finalNodesPool[i], finalNodesPool[j] = finalNodesPool[j], finalNodesPool[i]
	})
	fmt.Println(finalNodesPool)
	fmt.Println("开始挖矿...")
}

func getMineNodeAddress() string {
	count++
	return finalNodesPool[count%3].address
}
func vote() {
	for _, v := range voteNodesPool {
		superNodesPool[rand.Intn(len(superNodesPool))].votes += v.tokens
	}
}
func max() []int {
	var maxlist []int
	max := 0
	maxnum := 0
	for i := 0; i < 10; i++ {
		if max < superNodesPool[i].votes {
			max = superNodesPool[i].votes
			maxnum = i
		}
	}
	maxlist = append(maxlist, maxnum)
	max = 0
	maxnum = 0
	for i := 0; i < 10; i++ {
		if maxlist[0] == i {
			continue
		}
		if max < superNodesPool[i].votes {
			max = superNodesPool[i].votes
			maxnum = i
		}
	}
	maxlist = append(maxlist, maxnum)
	max = 0
	maxnum = 0
	for i := 0; i < 10; i++ {
		if maxlist[0] == i || maxlist[1] == i {
			continue
		}
		if max < superNodesPool[i].votes {
			max = superNodesPool[i].votes
			maxnum = i
		}
	}
	maxlist = append(maxlist, maxnum)
	return maxlist
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
