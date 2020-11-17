package blockchain

import (
	"flag"
	"go.etcd.io/bbolt"
	"os"
)

type Storage struct {
	db *bbolt.DB
}

func (s Storage) Append(value []byte, genesis func() *Block) []byte {
	s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		if b == nil {
			g := genesis()
			ser, err := g.Serialize()
			if err != nil {
				return err
			}
			err = b.Put([]byte(g.Hash), ser)
			if err != nil {
				return err
			}
			err = b.Put([]byte(blockKey), value)
			if err != nil {
				return err
			}
			return nil

		}
		return nil
	})
	return nil
}

func (s Storage) Tail() ([]byte, error) {
	var result []byte
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		result = b.Get([]byte(blockKey))
		return nil
	})
	return result, err
}

var (
	dbFilePath  string
	blockBucket = "block_chain_bucket"
	blockKey    = "_"
)

func NewStorage() *Storage {
	db, err := bbolt.Open(dbFilePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	return &Storage{db: db}
}

func init() {
	flag.StringVar(&dbFilePath, "db-file-path", os.TempDir()+"chain.db", "db file path")
}
