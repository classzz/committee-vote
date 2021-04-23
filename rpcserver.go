package main

import (
	"github.com/classzz/committee-vote/storage"
	"math/big"
)

type RPCService struct {
	RawDB *storage.RawDB
}

func (ser *RPCService) CurrentHeight(reply *big.Int) error {
	var err error
	if reply, err = ser.RawDB.CurrentHeight(); err != nil {
		return err
	}
	return nil
}
