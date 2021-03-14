package scanning

import (
	"fmt"
	"github.com/classzz/classzz/blockchain"
	"github.com/classzz/classzz/chaincfg"
	"github.com/classzz/classzz/cross"
	"github.com/classzz/classzz/rpcclient"
	"github.com/classzz/committee-vote/chains/ethereum"
	"github.com/classzz/committee-vote/chains/heco"
	"github.com/classzz/committee-vote/storage"
	"log"
	"math/big"
	"strconv"
	"time"
)

var (
	startInterval = 1 * time.Second
)

type Scanning struct {
	NodeClient  *rpcclient.Client
	MysqlClient *storage.MysqlClient
	EthClient   *ethereum.EthClient
	HecoClient  *heco.HecoClient
	MaxHeight   int64
	CloseCh     chan struct{}
}

func NewScanning(cfg *Config, client *storage.MysqlClient) *Scanning {

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
		log.Println(err)
	}

	scanning := &Scanning{
		NodeClient:  rpcClient,
		MysqlClient: client,
		CloseCh:     make(chan struct{}),
	}
	return scanning
}

func (s *Scanning) Start() {

	s.MaxHeight = s.MysqlClient.BlockFindMaxHeight()
	if s.MaxHeight != 0 {
		s.MaxHeight = s.MaxHeight + 1
	}

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
	s.MysqlClient.Stop()
}

func (s *Scanning) start() {

	// Get the latest height
	//blockCount := int64(50)
	blockCount, err := s.NodeClient.GetBlockCount()
	if err != nil {
		log.Println(err)
		return
	}

	// Get the maximum

	fmt.Println("blockCount height", s.MaxHeight)
	for ; s.MaxHeight < blockCount; s.MaxHeight++ {

		blockHash, err := s.NodeClient.GetBlockHash(s.MaxHeight)
		if err != nil {
			log.Println(err)
			return
		}

		block, err := s.NodeClient.GetBlock(blockHash.String())
		header := block.Header
		params := &chaincfg.MainNetParams
		dblock := &storage.CzzBlocks{
			PreviousBlockHash: header.PrevBlock.String(),
			Hash:              block.BlockHash().String(),
			Height:            s.MaxHeight,
			Version:           header.Version,
			VersionHex:        fmt.Sprintf("%08x", header.Version),
			Merkleroot:        header.MerkleRoot.String(),
			CTime:             header.Timestamp.Unix(),
			Nonce:             header.Nonce,
			Bits:              strconv.FormatInt(int64(header.Bits), 16),
			Difficulty:        getDifficultyRatio(header.Bits, params),
			IsMain:            1,
		}

		if e := s.MysqlClient.BlockInstall(dblock); e == 0 {
			if err := s.ProcessConvert(); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (s *Scanning) ProcessConvert() error {
	convs, err := s.NodeClient.GetConvertItems(nil, nil)
	if err != nil {
		return err
	}

	for _, conv := range convs {
		if s.MysqlClient.FindConvertItem(conv.MID) == nil {
			s.MysqlClient.ConvertItemInstall(conv)
			if conv.ConvertType == cross.ExpandedTxConvert_ECzz {
				if _, err := s.EthClient.Casting(conv); err != nil {
					return err
				}
			} else if conv.ConvertType == cross.ExpandedTxConvert_HCzz {
				if _, err := s.HecoClient.Casting(conv); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// getDifficultyRatio returns the proof-of-work difficulty as a multiple of the
// minimum difficulty using the passed bits field from the header of a block.
func getDifficultyRatio(bits uint32, params *chaincfg.Params) float64 {
	// The minimum difficulty is the max possible proof-of-work limit bits
	// converted back to a number.  Note this is not the same as the proof of
	// work limit directly because the block difficulty is encoded in a block
	// with the compact form which loses precision.
	max := blockchain.CompactToBig(params.PowLimitBits)
	target := blockchain.CompactToBig(bits)

	difficulty := new(big.Rat).SetFrac(max, target)
	outString := difficulty.FloatString(8)
	diff, err := strconv.ParseFloat(outString, 64)
	if err != nil {
		return 0
	}
	return diff
}
