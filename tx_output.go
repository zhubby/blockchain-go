package blockchain

import (
	"bytes"
	"encoding/gob"
)

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

// Lock signs the output
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey checks if the output can be used by the owner of the pubkey
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

type TxOutputs []TxOutput

func (outs TxOutputs) Serialize() ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	return buff.Bytes(), err
}

func (outs *TxOutputs) Deserialize(data []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(outs)
}
