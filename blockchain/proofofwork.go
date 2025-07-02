package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	"github.com/devduart/blockGo/utils"
)

// o proof-of-work eh basicamente o trabalho que se teve para gerar um hash valido
// quanto mais complexo a regra maior o trabalho computacional para gerar o hash
// no caso, quanto menor for o target mais processamento sera necessario para gerar o hash
// pois o hash valido precisa ser menor que o target

// quanto maior, mais tempo demora para gerar o hash
// significa que um hash para ser valido os 2 primeiros bytes (16 bits) precisam ser zero
const targetBits = 16

// estrutura de dados do proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	// inicializa com 1
	target := big.NewInt(1)
	// shift left
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	// o laco vai se repetir ate que o hash calculado seja menor que target
	// indicando ser um hash valido
	// apenas o contador eh alterado para tentar produzir um hash valido
	// por isso do laco, nonce sera incrementado ate que para aquele conjunto de dados e o nonce
	// o hash produzido eh valido
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)

		// gerar o hash
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	// retorna o numero de iteracoes e o hash valido
	return nonce, hash[:]
}

// verifica se o hash eh valido
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

// seta os dados do bloco
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}
