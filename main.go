package main

import (
	"fmt"
	"strconv"

	"github.com/devduart/blockGo/blockchain"
)

func main() {
	// Cria uma nova blockchain com o bloco gênesis
	bc := blockchain.NewBlockchain()

	// Adiciona dois blocos novos com dados de transações simuladas
	bc = blockchain.AddBlock(bc, "Send 1 GoHorse to Alice")
	bc = blockchain.AddBlock(bc, "Send 2 GoHorse to Bob")

	// Itera sobre a blockchain e imprime os dados de cada bloco
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
		fmt.Printf("Nonce: %d\n\n", block.Nonce)

		pow := blockchain.NewProofOfWork(&block)
		fmt.Printf("Pow: %s \n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
