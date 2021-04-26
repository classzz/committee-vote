package scanning

import (
	"github.com/classzz/classzz/rpcclient"
	"github.com/classzz/committee-vote/chains/bsc"
	"github.com/classzz/committee-vote/chains/ethereum"
	"github.com/classzz/committee-vote/chains/heco"
	"github.com/classzz/committee-vote/storage"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

var startInterval = 1 * time.Second

type Scanning struct {
	NodeClient *rpcclient.Client
	RawDB      *storage.RawDB
	EthClient  *ethereum.EthClient
	HecoClient *heco.HecoClient
	BscClient  *bsc.BscClient
	MaxHeight  int64
	CloseCh    chan struct{}
}

func NewScanning(cfg *Config, rawdb *storage.RawDB) *Scanning {

	connCfg := &rpcclient.ConnConfig{
		Host:         cfg.RpcHost,
		Endpoint:     "ws",
		User:         cfg.RpcUsername,
		Pass:         cfg.RpcPassword,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	rpcClient, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Error("rpcclient.New", "err", err)
	}

	scanning := &Scanning{
		NodeClient: rpcClient,
		RawDB:      rawdb,
		CloseCh:    make(chan struct{}),
	}
	return scanning
}

func (s *Scanning) Start() {
	CurrentHeight, err := s.RawDB.CurrentHeight()
	if err != nil {
		s.RawDB.SetCurrentHeight(0)
		CurrentHeight = big.NewInt(0)
	}

	s.MaxHeight = CurrentHeight.Int64()
	startTicker := time.NewTicker(startInterval)
out:
	for {
		select {
		case <-startTicker.C:
			s.start()
		case <-s.CloseCh:
			break out
		default:
		}
	}
}

func (s *Scanning) Stop() {
	close(s.CloseCh)
	// close rpc
	s.NodeClient.Shutdown()
	s.RawDB.Stop()
}

func (s *Scanning) start() {

	// Get the latest height
	blockCount, err := s.NodeClient.GetBlockCount()
	if err != nil {
		log.Error("GetBlockCount", "err", err)
		return
	}

	// Get the maximum
	log.Info("blockCount", "height", s.MaxHeight)
	if s.MaxHeight < blockCount {
		blockHash, err := s.NodeClient.GetBlockHash(blockCount)
		if err != nil {
			log.Error("GetBlockHash", "err", err)
			return
		}

		if blockHash == nil {
			return
		}

		if err := s.ProcessConvert(); err != nil {
			log.Error("ProcessConvert", "err", err)
			return
		}

		s.MaxHeight = blockCount
		if err := s.RawDB.SetCurrentHeight(blockCount); err != nil {
			log.Error("SetCurrentHeight", "err", err)
			return
		}
	}
}

func (s *Scanning) ProcessConvert() error {
	convs, err := s.NodeClient.GetConvertItems(nil, nil)
	if err != nil {
		return err
	}

	for _, conv := range convs {
		if item, err := s.RawDB.GetConvertItem(conv.MID); item == nil {
			if err = s.RawDB.SetConvertItem(conv); err != nil {
				return err
			}
			var exttx *types.Transaction
			//if conv.ConvertType == cross.ExpandedTxConvert_ECzz {
			//	if exttx, err = s.EthClient.Casting(conv); err != nil {
			//		return err
			//	}
			//} else if conv.ConvertType == cross.ExpandedTxConvert_HCzz {
			//	if exttx, err = s.HecoClient.Casting(conv); err != nil {
			//		return err
			//	}
			//} else if conv.ConvertType == cross.ExpandedTxConvert_BCzz {
			//	if exttx, err = s.BscClient.Casting(conv); err != nil {
			//		return err
			//	}
			//}
			if exttx != nil {
				conv.ConfirmExtTxHash = exttx.Hash().Hex()
				if err := s.RawDB.SetConvertItem(conv); err != nil {
					return err
				}

				if err := s.RawDB.SetExtTx(exttx); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
