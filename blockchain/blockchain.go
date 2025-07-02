package blockchain

import "slices"

// Blockchain representa a cadeia de blocos.
// Aqui é uma lista imutável de blocos (por valor).
type Blockchain struct {
	Blocks []Block
}

// NewGenesisBlock cria o primeiro bloco da cadeia, conhecido como "Genesis Block".
// Não possui hash anterior.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain cria uma nova blockchain com apenas o bloco gênesis.
// É a porta de entrada da cadeia.
func NewBlockchain() Blockchain {
	return Blockchain{
		Blocks: []Block{*NewGenesisBlock()},
	}
}

// AddBlock adiciona um novo bloco ao final da blockchain.
// Essa função não altera o estado original: retorna uma nova instância da cadeia.
func AddBlock(bc Blockchain, data string) Blockchain {
	// Obtém o último bloco existente na cadeia
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	// Cria um novo bloco com os dados e o hash do último bloco
	newBlock := NewBlock(data, lastBlock.Hash)

	// Cria uma nova slice (imutável) contendo os blocos anteriores + o novo bloco
	newChain := slices.Clone(bc.Blocks)    // cópia dos blocos antigos
	newChain = append(newChain, *newBlock) // adiciona o novo bloco
	return Blockchain{Blocks: newChain}
}
