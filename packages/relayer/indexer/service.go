package indexer

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/cyberhorsey/errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer/contracts/bridge"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer/contracts/icrosschainsync"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer/contracts/wyzth_zkevml1"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer/contracts/tokenvault"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer/message"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer/proof"
)

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

type ethClient interface {
	ChainID(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
}

type Service struct {
	eventRepo relayer.EventRepository
	blockRepo relayer.BlockRepository
	ethClient ethClient
	destRPC   *rpc.Client

	processingBlockHeight uint64

	bridge     relayer.Bridge
	destBridge relayer.Bridge

	processor *message.Processor

	relayerAddr common.Address

	blockBatchSize      uint64
	numGoroutines       int
	subscriptionBackoff time.Duration

	wyzth_zkevml1 *wyzth_zkevml1.WYzth_zkevmL1
}

type NewServiceOpts struct {
	EventRepo                     relayer.EventRepository
	BlockRepo                     relayer.BlockRepository
	EthClient                     *ethclient.Client
	DestEthClient                 *ethclient.Client
	RPCClient                     *rpc.Client
	DestRPCClient                 *rpc.Client
	ECDSAKey                      string
	BridgeAddress                 common.Address
	DestBridgeAddress             common.Address
	SrcWYzth_zkevmAddress               common.Address
	DestWYzth_zkevmAddress              common.Address
	DestTokenVaultAddress         common.Address
	SrcSignalServiceAddress       common.Address
	BlockBatchSize                uint64
	NumGoroutines                 int
	SubscriptionBackoff           time.Duration
	Confirmations                 uint64
	ProfitableOnly                relayer.ProfitableOnly
	HeaderSyncIntervalInSeconds   int64
	ConfirmationsTimeoutInSeconds int64
}

func NewService(opts NewServiceOpts) (*Service, error) {
	if opts.EventRepo == nil {
		return nil, relayer.ErrNoEventRepository
	}

	if opts.BlockRepo == nil {
		return nil, relayer.ErrNoBlockRepository
	}

	if opts.EthClient == nil {
		return nil, relayer.ErrNoEthClient
	}

	if opts.ECDSAKey == "" {
		return nil, relayer.ErrNoECDSAKey
	}

	if opts.DestEthClient == nil {
		return nil, relayer.ErrNoEthClient
	}

	if opts.BridgeAddress == ZeroAddress {
		return nil, relayer.ErrNoBridgeAddress
	}

	if opts.DestBridgeAddress == ZeroAddress {
		return nil, relayer.ErrNoBridgeAddress
	}

	if opts.RPCClient == nil {
		return nil, relayer.ErrNoRPCClient
	}

	privateKey, err := crypto.HexToECDSA(opts.ECDSAKey)
	if err != nil {
		return nil, errors.Wrap(err, "crypto.HexToECDSA")
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.Wrap(err, "publicKey.(*ecdsa.PublicKey)")
	}

	relayerAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	srcBridge, err := bridge.NewBridge(opts.BridgeAddress, opts.EthClient)
	if err != nil {
		return nil, errors.Wrap(err, "bridge.NewBridge")
	}

	destBridge, err := bridge.NewBridge(opts.DestBridgeAddress, opts.DestEthClient)
	if err != nil {
		return nil, errors.Wrap(err, "bridge.NewBridge")
	}

	prover, err := proof.New(opts.EthClient)
	if err != nil {
		return nil, errors.Wrap(err, "proof.New")
	}

	destHeaderSyncer, err := icrosschainsync.NewICrossChainSync(opts.DestWYzth_zkevmAddress, opts.DestEthClient)
	if err != nil {
		return nil, errors.Wrap(err, "icrosschainsync.NewWYzth_zkevmL2")
	}

	var wyzth_zkevmL1 *wyzth_zkevml1.WYzth_zkevmL1
	if opts.SrcWYzth_zkevmAddress != ZeroAddress {
		wyzth_zkevmL1, err = wyzth_zkevml1.NewWYzth_zkevmL1(opts.SrcWYzth_zkevmAddress, opts.EthClient)
		if err != nil {
			return nil, errors.Wrap(err, "wyzth_zkevml1.NewWYzth_zkevmL1")
		}
	}

	destTokenVault, err := tokenvault.NewTokenVault(opts.DestTokenVaultAddress, opts.DestEthClient)
	if err != nil {
		return nil, errors.Wrap(err, "tokenvault.NewTokenVault")
	}

	processor, err := message.NewProcessor(message.NewProcessorOpts{
		Prover:                        prover,
		ECDSAKey:                      privateKey,
		RPCClient:                     opts.RPCClient,
		DestETHClient:                 opts.DestEthClient,
		DestBridge:                    destBridge,
		EventRepo:                     opts.EventRepo,
		DestHeaderSyncer:              destHeaderSyncer,
		RelayerAddress:                relayerAddr,
		Confirmations:                 opts.Confirmations,
		SrcETHClient:                  opts.EthClient,
		ProfitableOnly:                opts.ProfitableOnly,
		HeaderSyncIntervalSeconds:     opts.HeaderSyncIntervalInSeconds,
		SrcSignalServiceAddress:       opts.SrcSignalServiceAddress,
		ConfirmationsTimeoutInSeconds: opts.ConfirmationsTimeoutInSeconds,
		DestTokenVault:                destTokenVault,
	})
	if err != nil {
		return nil, errors.Wrap(err, "message.NewProcessor")
	}

	return &Service{
		blockRepo: opts.BlockRepo,
		eventRepo: opts.EventRepo,
		ethClient: opts.EthClient,
		destRPC:   opts.DestRPCClient,

		bridge:     srcBridge,
		destBridge: destBridge,
		wyzth_zkevml1:    wyzth_zkevmL1,

		processor: processor,

		relayerAddr: relayerAddr,

		blockBatchSize:      opts.BlockBatchSize,
		numGoroutines:       opts.NumGoroutines,
		subscriptionBackoff: opts.SubscriptionBackoff,
	}, nil
}
