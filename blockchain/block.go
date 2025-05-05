package blockchain

import (
	"time"
)

// Block representa um bloco da blockchain.
// É imutável e contém dados, hash, hash do bloco anterior e um timestamp.
type Block struct {
	Timestamp     int64  // momento da criação do bloco
	Data          []byte // dados arbitrários (ex: transações)
	PrevBlockHash []byte // hash do bloco anterior
	Hash          []byte // hash do próprio bloco (calculado a partir dos campos acima)
	Nonce         int    // indica o numero de iteracoes para encontrar um hash valido
}

// NewBlock cria um novo bloco a partir dos dados e do hash do bloco anterior.
// Essa função é pura: dado o mesmo input, sempre gera o mesmo resultado.

// alterado para que seja passado os dados base para o ProofOfWork gerar um hash valido
// o hash e nonce é populado pelo ProofOfWork
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevHash,
		[]byte{},
		0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
