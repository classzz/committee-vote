package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
	"math/big"
)

var (
	currentHeight = "CurrentHeight"
)

type RawDB struct {
	DB *leveldb.DB
}

func (d RawDB) Stop() error {
	return d.DB.Close()
}

func (d RawDB) CurrentHeight() (*big.Int, error) {
	if data, err := d.DB.Get([]byte(currentHeight), nil); err != nil {
		return nil, err
	} else {
		return big.NewInt(0).SetBytes(data), nil
	}
}

func (d RawDB) SetCurrentHeight(height int64) error {
	heightB := big.NewInt(height).Bytes()
	if err := d.DB.Put([]byte(currentHeight), heightB, nil); err != nil {
		return err
	}
	return nil
}
