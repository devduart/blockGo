package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// Block representa um bloco da blockchain.
// É imutável e contém dados, hash, hash do bloco anterior e um timestamp.
type Block struct {
	Timestamp     int64  // momento da criação do bloco
	Data          []byte // dados arbitrários (ex: transações)
	PrevBlockHash []byte // hash do bloco anterior
	Hash          []byte // hash do próprio bloco (calculado a partir dos campos acima)
}

// NewBlock cria um novo bloco a partir dos dados e do hash do bloco anterior.
// Essa função é pura: dado o mesmo input, sempre gera o mesmo resultado.
func NewBlock(data string, prevHash []byte) Block {
	timestamp := time.Now().Unix()
	// headers é a concatenação dos dados usados para gerar o hash
	headers := bytes.Join([][]byte{prevHash, []byte(data), []byte(string(timestamp))}, []byte{})
	hash := sha256.Sum256(headers)

	return Block{
		Timestamp:     timestamp,
		Data:          []byte(data),
		PrevBlockHash: prevHash,
		Hash:          hash[:],
	}
}
