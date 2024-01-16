package main

import (
	"bytes"
	"crypto/sha256"

	// "fmt"
	"math/big"
)

const targetBits = 20

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			pow.block.Data,
			IntToHex(pow.block.Time),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	// var hashInt big.Int //hash的整形表示
	var hash []byte
	nonce := 0
	// fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	/*
		实现Hashcash算法：对nonce从0开始进行遍历，计算每一次哈希是否满足条件
		可能会用到的包及函数：big.Int.Cmp(),big.Int.SetBytes()
	*/
	// temp := big.NewInt(1)
	// hashInt := big.NewInt(2)
	// for hasher := sha256.New(); hashInt.Cmp(temp) != 0; nonce++ {
	// 	hasher.Reset()
	// 	hasher.Write(pow.prepareData(nonce))
	// 	hash = hasher.Sum(nil)
	// 	hashInt.SetBytes(hash)
	// 	temp.SetBytes(hash)
	// 	temp.AndNot(temp, pow.target)
	// }
	hasher := sha256.New()
	hashInt := big.NewInt(2)
	hasher.Reset()
	hasher.Write(pow.prepareData(nonce))
	hash = hasher.Sum(nil)
	hashInt.SetBytes(hash)
	for ; hashInt.Cmp(pow.target) != -1; nonce++ {
		hasher.Reset()
		hasher.Write(pow.prepareData(nonce))
		hash = hasher.Sum(nil)
		hashInt.SetBytes(hash)
	}
	// fmt.Printf("\r%x", hash)
	// fmt.Print("\n\n")
	return nonce, hash[:]

}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	temp := big.NewInt(1)
	hashInt.SetBytes(pow.block.Hash)
	temp.SetBytes(pow.block.Hash)
	temp.AndNot(temp, pow.target)
	return hashInt.Cmp(temp) == 0
}
