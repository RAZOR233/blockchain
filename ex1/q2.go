package main

import (
	"base58"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

func hash160(input []byte) []byte {
	hasher := sha256.New()
	hasher.Write(input)
	out1 := hasher.Sum([]byte(""))
	hasher2 := ripemd160.New()
	hasher2.Write(out1)
	out := hasher2.Sum([]byte(""))
	return out
}

func hash256(input []byte) []byte {
	hasher := sha256.New()
	hasher.Write(input)
	out1 := hasher.Sum([]byte(""))
	hasher2 := sha256.New()
	hasher2.Write(out1)
	out2 := hasher2.Sum([]byte(""))
	return out2
}

func address(public_key string) string {
	version_byte := []byte{0x6f}

	input, _ := hex.DecodeString(public_key)
	FingerPrint := hash160(input)
	t := append(version_byte, FingerPrint...)
	// fmt.Println(t)
	// fmt.Print(t[:8])
	CheckSum := hash256(t)[:4]
	// fmt.Println(t2)
	// t3 := version_byte + FingerPrint + CheckSum
	t3 := append(t, CheckSum...)
	// fmt.Printf("FingerPrint:%s\n", hex.EncodeToString(FingerPrint))
	// fmt.Printf("CheckSum:%s\n", hex.EncodeToString(CheckSum))
	Address := base58.Encode([]byte(t3), base58.BitcoinAlphabet)
	// fmt.Printf("%s", Address)
	return Address
}
func main() {
	public_key1 := "02b1ebcdbac723f7444fdfb8e83b13bd14fe679c59673a519df6a1038c07b719c6"
	public_key2 := "036e69a3e7c303935403d5b96c47b7c4fa8a80ca569735284a91d930f0f49afa86"
	fmt.Printf("the address of public key1 is:%s\n", address(public_key1))
	fmt.Printf("the address of public key2 is:%s\n", address(public_key2))
}
