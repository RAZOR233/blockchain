package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

const targetBits = 16

type ProofOfWork struct {
	block  *block
	target *big.Int
}

func NewProofOfWork(b *block) *ProofOfWork {
	count := 0
	for i := 0; i < targetBits; i++ {
		count = count << 1
		count += 1
	}
	tmp := big.NewInt(int64(count))
	//tmp.Lsh(tmp, uint(targetBits))
	lasthash, _ := hex.DecodeString(b.Lasthash)
	Int := big.NewInt(2)
	Int.SetBytes(lasthash)
	target := big.NewInt(1)
	target.And(tmp, Int)
	target.Lsh(target, uint(256-targetBits))
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
	count := 0
	for i := 0; i < targetBits; i++ {
		count = count << 1
		count += 1
	}
	tmp := big.NewInt(int64(count))
	tmp.Lsh(tmp, uint(256-targetBits))

	hasher := sha256.New()
	hashInt := big.NewInt(2)
	hasher.Reset()
	hasher.Write(pow.prepareData(nonce))
	hash = hasher.Sum(nil)
	hashInt.SetBytes(hash)
	hashInt.And(hashInt, tmp)
	for ; hashInt.Cmp(pow.target) != 0; nonce++ {
		hasher.Reset()
		hasher.Write(pow.prepareData(nonce))
		hash = hasher.Sum(nil)
		hashInt.SetBytes(hash)
		hashInt.And(hashInt, tmp)
	}
	fmt.Printf("\r%x", hash)
	fmt.Print("\n\n")
	return nonce, hex.EncodeToString(hash[:])

}
