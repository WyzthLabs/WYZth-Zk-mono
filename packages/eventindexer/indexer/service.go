package indexer

import (
	"time"

	"github.com/cyberhorsey/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/eventindexer"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/eventindexer/contracts/wyzth_zkevml1"
)

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

type Service struct {
	eventRepo eventindexer.EventRepository
	blockRepo eventindexer.BlockRepository
	statRepo  eventindexer.StatRepository
	ethClient *ethclient.Client

	processingBlockHeight uint64

	blockBatchSize      uint64
	subscriptionBackoff time.Duration

	wyzth_zkevml1 *wyzth_zkevml1.WYzth_zkevmL1
}

type NewServiceOpts struct {
	EventRepo           eventindexer.EventRepository
	BlockRepo           eventindexer.BlockRepository
	StatRepo            eventindexer.StatRepository
	EthClient           *ethclient.Client
	RPCClient           *rpc.Client
	SrcWYzth_zkevmAddress     common.Address
	BlockBatchSize      uint64
	SubscriptionBackoff time.Duration
}

func NewService(opts NewServiceOpts) (*Service, error) {
	if opts.EventRepo == nil {
		return nil, eventindexer.ErrNoEventRepository
	}

	if opts.EthClient == nil {
		return nil, eventindexer.ErrNoEthClient
	}

	if opts.RPCClient == nil {
		return nil, eventindexer.ErrNoRPCClient
	}

	wyzth_zkevmL1, err := wyzth_zkevml1.NewWYzth_zkevmL1(opts.SrcWYzth_zkevmAddress, opts.EthClient)
	if err != nil {
		return nil, errors.Wrap(err, "contracts.NewWYzth_zkevmL1")
	}

	return &Service{
		eventRepo: opts.EventRepo,
		blockRepo: opts.BlockRepo,
		statRepo:  opts.StatRepo,
		ethClient: opts.EthClient,
		wyzth_zkevml1:   wyzth_zkevmL1,

		blockBatchSize:      opts.BlockBatchSize,
		subscriptionBackoff: opts.SubscriptionBackoff,
	}, nil
}
