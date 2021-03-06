package storage

import (
	"github.com/classzz/classzz/btcjson"
	"github.com/classzz/classzz/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
)

var (
	currentHeight = "CurrentHeight"
	Mid           = "mid-"
	Ext           = "ext-"
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

func (d RawDB) GetConvertItem(mid *big.Int) (*btcjson.ConvertItemsResult, error) {
	if data, err := d.DB.Get([]byte(Mid+mid.String()), nil); err != nil {
		return nil, err
	} else {
		item := &btcjson.ConvertItemsResult{}
		if err := rlp.DecodeBytes(data, item); err != nil {
			return nil, err
		}
		return item, nil
	}
}

func (d RawDB) GetConvertItemAll() ([]*btcjson.ConvertItemsResult, error) {

	iter := d.DB.NewIterator(util.BytesPrefix([]byte(Mid)), nil)
	list := make([]*btcjson.ConvertItemsResult, 0, 0)
	for iter.Next() {
		value := iter.Value()
		item := &btcjson.ConvertItemsResult{}
		if err := rlp.DecodeBytes(value, item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		return nil, err
	}
	return list, nil
}

func (d RawDB) SetConvertItem(item *btcjson.ConvertItemsResult) error {
	key := []byte(Mid + item.MID.String())
	value, err := rlp.EncodeToBytes(item)
	if err != nil {
		return err
	}
	if err := d.DB.Put(key, value, nil); err != nil {
		return err
	}
	return nil
}

func (d RawDB) SetExtTx(tx *types.Transaction) error {
	key := []byte(Ext + tx.Hash().String())
	value, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	if err := d.DB.Put(key, value, nil); err != nil {
		return err
	}
	return nil
}
