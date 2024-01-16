package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

const targetBits = 12

type ProofOfWork struct {
	block  *block
	target *big.Int
}

func NewProofOfWork(b *block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	//fmt.Print(hex.EncodeToString(target.Bytes()))
	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	lasthash, _ := hex.DecodeString(pow.block.Lasthash)
	data := bytes.Join(

		[][]byte{
			lasthash,
			[]byte(pow.block.Data),
			[]byte(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, string) {
	// var hashInt big.Int //hash的整形表示
	var hash []byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

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
	fmt.Printf("\r%x", hash)
	fmt.Print("\n\n")
	return nonce, hex.EncodeToString(hash[:])

}

// func (pow *ProofOfWork) Validate() bool {
// 	var hashInt big.Int
// 	temp := big.NewInt(1)
// 	hashInt.SetBytes(pow.block.Hash)
// 	temp.SetBytes(pow.block.Hash)
// 	temp.AndNot(temp, pow.target)
// 	return hashInt.Cmp(temp) == 0
// }
