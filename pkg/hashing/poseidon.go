package hashing

import (
	"errors"
	"math/big"

	"github.com/iden3/go-iden3-crypto/v2/poseidon"
)

// spongeChunkSize defines the size of each chunk in bytes for processing data in the sponge hashing algorithm.
const spongeChunkSize = 31

// spongeInputs defines the default number of input blocks for the sponge hashing algorithm, limited between 2 and 16.
const spongeInputs = 16

// HashBytes returns a sponge hash of a msg byte slice split into blocks of 31 bytes.
//
// Source: https://github.com/iden3/go-iden3-crypto/blob/d59dca822c2ad154e1a9011514e2c01c6581597b/poseidon/poseidon.go#L160
//
// Copied verbatim. This method was removed in v2.
func HashBytes(msg []byte) (*big.Int, error) {
	return HashBytesX(msg, spongeInputs)
}

// HashBytesX returns a sponge hash of a msg byte slice split into blocks of 31 bytes.
//
// Source: https://github.com/iden3/go-iden3-crypto/blob/d59dca822c2ad154e1a9011514e2c01c6581597b/poseidon/poseidon.go#L165
//
// Copied verbatim. This method was removed in v2.
func HashBytesX(msg []byte, frameSize int) (*big.Int, error) {
	if frameSize < 2 || frameSize > 16 {
		return nil, errors.New("incorrect frame size")
	}

	// not used inputs default to zero
	inputs := make([]*big.Int, frameSize)
	for j := 0; j < frameSize; j++ {
		inputs[j] = new(big.Int)
	}
	dirty := false
	var hash *big.Int
	var err error

	k := 0
	for i := 0; i < len(msg)/spongeChunkSize; i++ {
		dirty = true
		inputs[k].SetBytes(msg[spongeChunkSize*i : spongeChunkSize*(i+1)])
		if k == frameSize-1 {
			hash, err = poseidon.Hash(inputs)
			dirty = false
			if err != nil {
				return nil, err
			}
			inputs = make([]*big.Int, frameSize)
			inputs[0] = hash
			for j := 1; j < frameSize; j++ {
				inputs[j] = new(big.Int)
			}
			k = 1
		} else {
			k++
		}
	}

	if len(msg)%spongeChunkSize != 0 {
		// the last chunk of the message is less than 31 bytes
		// zero padding it, so that 0xdeadbeaf becomes
		// 0xdeadbeaf000000000000000000000000000000000000000000000000000000
		var buf [spongeChunkSize]byte
		copy(buf[:], msg[(len(msg)/spongeChunkSize)*spongeChunkSize:])
		inputs[k] = new(big.Int).SetBytes(buf[:])
		dirty = true
	}

	if dirty {
		// we haven't hashed something in the main sponge loop and need to do hash here
		hash, err = poseidon.Hash(inputs)
		if err != nil {
			return nil, err
		}
	}

	return hash, nil
}
