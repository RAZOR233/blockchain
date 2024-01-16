package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type tree struct {
	hash  []byte
	left  *tree
	right *tree
	depth int
	num   int
}

var a int

func hash(input []byte) []byte {
	hasher := sha256.New()
	hasher.Write(input)
	out1 := hasher.Sum([]byte(""))
	return out1
}

// 创建节点
func create_node(hash []byte, n int) *tree {
	return &tree{hash, nil, nil, n, 0}
}
func create_tree(n int, p *tree, list [16]string) bool {
	cur := p
	if cur.depth == n {
		k, _ := hex.DecodeString(list[a])
		cur.hash = hash(k)
		a++
		cur.num = a
		return true
	} else {
		k2, _ := hex.DecodeString("")
		cur.left = create_node(k2, cur.depth+1)
		cur.right = create_node(k2, cur.depth+1)
		create_tree(n, cur.left, list)
		create_tree(n, cur.right, list)
		return true
	}
}
func see_tree(p *tree, n int) {
	if p != nil {
		fmt.Println(p.num, hex.EncodeToString(p.hash))
		see_tree(p.left, n)
		see_tree(p.right, n)
	}

	// fmt.Printf("%d,%x\n", p.depth, p.hash)
}
func work(p *tree) {
	if hex.EncodeToString(p.left.hash) != "" {
		p.hash = hash(append(p.left.hash, p.right.hash...))
	} else {
		work(p.left)
		work(p.right)
	}
}

func compareMerkleTree(tree1 *tree, tree2 *tree) int {
	var i int
	var p1 *tree
	var p2 *tree
	p1 = tree1
	p2 = tree2
	for i = 0; i <= 3; i++ {

		// fmt.Println(hex.EncodeToString(p1.hash), " ", hex.EncodeToString(p2.hash))
		if hex.EncodeToString(p1.left.hash) != hex.EncodeToString(p2.left.hash) {
			p1 = p1.left
			p2 = p2.left
			// fmt.Println(p1.num)
		} else {
			p1 = p1.right
			p2 = p2.right
			// fmt.Println(p1.num)
		}
	}
	return (p1.num)
}

// 插入节点
func main() {
	a = 0
	var i int
	var list1 = [16]string{"20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35"}
	k, _ := hex.DecodeString("")
	var b = create_node(k, 1)
	create_tree(5, b, list1)

	for i = 0; i <= 4; i++ {
		work(b)
	}
	see_tree(b, 5)
	fmt.Print("\n")
	a = 0
	var list2 = [16]string{"20", "21", "22", "23", "24", "25", "26", "257", "28", "29", "30", "31", "32", "33", "34", "35"}
	var c = create_node(k, 1)
	create_tree(5, c, list2)
	for i = 0; i <= 4; i++ {
		work(c)
	}
	see_tree(c, 5)
	fmt.Printf("发生错误的节点为：%d", compareMerkleTree(b, c))
}
