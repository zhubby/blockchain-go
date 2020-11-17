package blockchain

import (
	"go.etcd.io/bbolt"
	"log"
)

type ChainIterator struct {
	currentHash []byte
	db          *bbolt.DB
}

func (i *ChainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		encodedBlock := b.Get(i.currentHash)
		return block.Deserialize(encodedBlock)
	})
	if err != nil {
		log.Fatal(err)
	}
	i.currentHash = block.ParentBlockHash
	return block
}
