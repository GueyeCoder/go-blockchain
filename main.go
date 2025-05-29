package main

import (
	"fmt"
	"log"
	"strings"

	"goblockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKey())
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(strings.Repeat("*", 25))
	fmt.Println(w.PublicKey())
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println(w.PublicKeyStr())
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println(w.BlockchainAddress())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	fmt.Printf("signature: %s\n", t.GenerateSignature())

}
