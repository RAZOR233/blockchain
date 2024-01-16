// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	t := time.Now()

// 	block := NewBlock("Genesis Block", []byte{})

// 	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
// 	fmt.Printf("Time: %s\n", time.Unix(block.Time, 0).Format("2006-01-02 15:04:05"))
// 	fmt.Printf("Data: %s\n", block.Data)
// 	fmt.Printf("Hash: %x\n", block.Hash)

// 		fmt.Println("Time using: ", time.Since(t))
// 	}

// ex2
// package main

// import (
// 	"fmt"
// )

//	func main() {
//		bc := NewBlockchain()
//		bc.AddBlock("Send 1 BTC to Ivan")
//		bc.AddBlock("Send 2 more BTC to Ivan")
//		for _, block := range bc.blocks {
//			fmt.Printf("PrevHash: %x\n", block.PrevHash)
//			fmt.Printf("Data: %s\n", block.Data)
//			fmt.Printf("Hash: %x\n", block.Hash)
//			fmt.Println()
//		}
//	}
// go build .\main.go .\block.go .\blockchain.go .\proofofWork.go .\utils.go .\datasave.go
// go build .\main.go .\block.go .\blockchain.go .\proofofWork.go .\utils.go
// .\main.exe
// ex3
// package main

// import (
// 	"fmt"
// )

//	func main() {
//		bc := NewBlockchain()
//		bc.AddBlock("Send 1 BTC to Ivan")
//		bc.AddBlock("Send 2 more BTC to Ivan")
//
//	for _, block := range bc.blocks {
//		fmt.Printf("PrevHash: %x\n", block.PrevHash)
//		fmt.Printf("Data: %s\n", block.Data)
//		fmt.Printf("Hash: %x\n", block.Hash)
//		fmt.Printf("POW: %t\n", NewProofOfWork(block).Validate())
//		fmt.Println()
//	}
//
//	}
//
// exex
package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var AllCode2 []recode
var AllCode3 []recode

func block_to_code(bbc *Blockchain) {

	for _, block := range bbc.blocks {
		t := recode{len(AllCode2) + 1, hex.EncodeToString(block.PrevHash), string(block.Data), hex.EncodeToString(block.Hash), strconv.Itoa(block.nouce)}
		AllCode2 = append(AllCode2, t)
	}
	AllCode = AllCode2
}
func code_to_block(bbc *Blockchain) *Blockchain {
	for _, code := range AllCode {
		a, _ := strconv.Atoi(code.nouce)
		prehash, _ := hex.DecodeString(code.PrevHash)
		hash, _ := hex.DecodeString(code.Hash)
		t := &Block{0, []byte(code.Data), prehash, hash, a}
		bbc.blocks = append(bbc.blocks, t)
	}
	bbc.blocks = bbc.blocks[1:]
	return bbc
}
func output_block(bbc *Blockchain) {
	for _, block := range bbc.blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("POW: %t\n", NewProofOfWork(block).Validate())
		fmt.Println()
	}
}
func main() {
	bc := NewBlockchain()
	newblock := flag.NewFlagSet("newblock", flag.ExitOnError)
	if os.Args[1] == "listblocks" {
		GetData()
		if len(AllCode) == 0 {
			fmt.Println("No existing block found!Creating a new one...")
			fmt.Printf("Mining the block containing \"%s\"\n", "Genesis Block")
			output_block(bc)
			block_to_code(bc)
			SaveData()
		} else {
			bc = code_to_block(bc)
			output_block(bc)
			block_to_code(bc)
			SaveData()
		}

	}
	if os.Args[1] == "newblock" {
		GetData()
		if len(AllCode) == 0 {
			data := newblock.String("data", "", "Input data")
			newblock.Parse(os.Args[2:])
			bc.AddBlock(*data)
			bc.blocks = bc.blocks[1:]
			fmt.Printf("Mining the block containing \"%s\"\n", *data)
			block_to_code(bc)
			SaveData()
			fmt.Println("success!")
		} else {
			data := newblock.String("data", "", "Input data")
			newblock.Parse(os.Args[2:])
			bc = code_to_block(bc)
			fmt.Printf("Mining the block containing \"%s\"\n", *data)
			bc.AddBlock(*data)
			block_to_code(bc)
			SaveData()
			fmt.Println("success!")
		}

	}
}

// GetData()
// bc = code_to_block(bc)
// output_block(bc)
// bc.AddBlock("hi")
// block_to_code(bc)
// SaveData()
