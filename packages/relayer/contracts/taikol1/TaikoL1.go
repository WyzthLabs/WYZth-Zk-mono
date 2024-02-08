// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wyzth_zkevml1

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// WYzth_zkevmDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type WYzth_zkevmDataBlockMetadata struct {
	Id                uint64
	Timestamp         uint64
	L1Height          uint64
	L1Hash            [32]byte
	MixHash           [32]byte
	DepositsRoot      [32]byte
	TxListHash        [32]byte
	TxListByteStart   *big.Int
	TxListByteEnd     *big.Int
	GasLimit          uint32
	Beneficiary       common.Address
	CacheTxListInfo   uint8
	Treasure          common.Address
	DepositsProcessed []WYzth_zkevmDataEthDeposit
}

// WYzth_zkevmDataConfig is an auto generated low-level Go binding around an user-defined struct.
type WYzth_zkevmDataConfig struct {
	ChainId                   *big.Int
	MaxNumProposedBlocks      *big.Int
	RingBufferSize            *big.Int
	MaxVerificationsPerTx     *big.Int
	BlockMaxGasLimit          *big.Int
	MaxTransactionsPerBlock   *big.Int
	MaxBytesPerTxList         *big.Int
	MinTxGasLimit             *big.Int
	TxListCacheExpiry         *big.Int
	ProofCooldownPeriod       *big.Int
	SystemProofCooldownPeriod *big.Int
	RealProofSkipSize         *big.Int
	EthDepositGas             *big.Int
	EthDepositMaxFee          *big.Int
	MinEthDepositsPerBlock    uint64
	MaxEthDepositsPerBlock    uint64
	MaxEthDepositAmount       *big.Int
	MinEthDepositAmount       *big.Int
	ProofTimeTarget           uint64
	AdjustmentQuotient        uint8
	RelaySignalRoot           bool
}

// WYzth_zkevmDataEthDeposit is an auto generated low-level Go binding around an user-defined struct.
type WYzth_zkevmDataEthDeposit struct {
	Recipient common.Address
	Amount    *big.Int
}

// WYzth_zkevmDataForkChoice is an auto generated low-level Go binding around an user-defined struct.
type WYzth_zkevmDataForkChoice struct {
	Key        [32]byte
	BlockHash  [32]byte
	SignalRoot [32]byte
	ProvenAt   uint64
	Prover     common.Address
	GasUsed    uint32
}

// WYzth_zkevmDataStateVariables is an auto generated low-level Go binding around an user-defined struct.
type WYzth_zkevmDataStateVariables struct {
	BlockFee                uint64
	AccBlockFees            uint64
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	NumBlocks               uint64
	ProofTimeIssued         uint64
	LastVerifiedBlockId     uint64
	AccProposedAt           uint64
	NextEthDepositToProcess uint64
	NumEthDeposits          uint64
}

// WYzth_zkevmL1MetaData contains all meta data concerning the WYzth_zkevmL1 contract.
var WYzth_zkevmL1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_PROHIBITED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_PROHIBITED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"depositsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"cacheTxListInfo\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"treasure\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structWYzth_zkevmData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structWYzth_zkevmData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structWYzth_zkevmData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEtherToL2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositWYzth_zkevmToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_proposedAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumProposedBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ringBufferSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txListCacheExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"systemProofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"realProofSkipSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"ethDepositMinCountPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ethDepositMaxCountPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"ethDepositMaxAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"ethDepositMinAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"adjustmentQuotient\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"relaySignalRoot\",\"type\":\"bool\"}],\"internalType\":\"structWYzth_zkevmData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"getForkChoice\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasUsed\",\"type\":\"uint32\"}],\"internalType\":\"structWYzth_zkevmData.ForkChoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposedAt\",\"type\":\"uint64\"}],\"name\":\"getProofReward\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEthDeposits\",\"type\":\"uint64\"}],\"internalType\":\"structWYzth_zkevmData.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getWYzth_zkevmTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getVerifierName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_initBlockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initProofTimeIssued\",\"type\":\"uint64\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"}],\"name\":\"proposeBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"depositsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"cacheTxListInfo\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"treasure\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structWYzth_zkevmData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"internalType\":\"structWYzth_zkevmData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved71\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved72\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved91\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawWYzth_zkevmToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WYzth_zkevmL1ABI is the input ABI used to generate the binding from.
// Deprecated: Use WYzth_zkevmL1MetaData.ABI instead.
var WYzth_zkevmL1ABI = WYzth_zkevmL1MetaData.ABI

// WYzth_zkevmL1 is an auto generated Go binding around an Ethereum contract.
type WYzth_zkevmL1 struct {
	WYzth_zkevmL1Caller     // Read-only binding to the contract
	WYzth_zkevmL1Transactor // Write-only binding to the contract
	WYzth_zkevmL1Filterer   // Log filterer for contract events
}

// WYzth_zkevmL1Caller is an auto generated read-only Go binding around an Ethereum contract.
type WYzth_zkevmL1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WYzth_zkevmL1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WYzth_zkevmL1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WYzth_zkevmL1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WYzth_zkevmL1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WYzth_zkevmL1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WYzth_zkevmL1Session struct {
	Contract     *WYzth_zkevmL1          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WYzth_zkevmL1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WYzth_zkevmL1CallerSession struct {
	Contract *WYzth_zkevmL1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WYzth_zkevmL1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WYzth_zkevmL1TransactorSession struct {
	Contract     *WYzth_zkevmL1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WYzth_zkevmL1Raw is an auto generated low-level Go binding around an Ethereum contract.
type WYzth_zkevmL1Raw struct {
	Contract *WYzth_zkevmL1 // Generic contract binding to access the raw methods on
}

// WYzth_zkevmL1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WYzth_zkevmL1CallerRaw struct {
	Contract *WYzth_zkevmL1Caller // Generic read-only contract binding to access the raw methods on
}

// WYzth_zkevmL1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WYzth_zkevmL1TransactorRaw struct {
	Contract *WYzth_zkevmL1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWYzth_zkevmL1 creates a new instance of WYzth_zkevmL1, bound to a specific deployed contract.
func NewWYzth_zkevmL1(address common.Address, backend bind.ContractBackend) (*WYzth_zkevmL1, error) {
	contract, err := bindWYzth_zkevmL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1{WYzth_zkevmL1Caller: WYzth_zkevmL1Caller{contract: contract}, WYzth_zkevmL1Transactor: WYzth_zkevmL1Transactor{contract: contract}, WYzth_zkevmL1Filterer: WYzth_zkevmL1Filterer{contract: contract}}, nil
}

// NewWYzth_zkevmL1Caller creates a new read-only instance of WYzth_zkevmL1, bound to a specific deployed contract.
func NewWYzth_zkevmL1Caller(address common.Address, caller bind.ContractCaller) (*WYzth_zkevmL1Caller, error) {
	contract, err := bindWYzth_zkevmL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1Caller{contract: contract}, nil
}

// NewWYzth_zkevmL1Transactor creates a new write-only instance of WYzth_zkevmL1, bound to a specific deployed contract.
func NewWYzth_zkevmL1Transactor(address common.Address, transactor bind.ContractTransactor) (*WYzth_zkevmL1Transactor, error) {
	contract, err := bindWYzth_zkevmL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1Transactor{contract: contract}, nil
}

// NewWYzth_zkevmL1Filterer creates a new log filterer instance of WYzth_zkevmL1, bound to a specific deployed contract.
func NewWYzth_zkevmL1Filterer(address common.Address, filterer bind.ContractFilterer) (*WYzth_zkevmL1Filterer, error) {
	contract, err := bindWYzth_zkevmL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1Filterer{contract: contract}, nil
}

// bindWYzth_zkevmL1 binds a generic wrapper to an already deployed contract.
func bindWYzth_zkevmL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WYzth_zkevmL1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WYzth_zkevmL1 *WYzth_zkevmL1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WYzth_zkevmL1.Contract.WYzth_zkevmL1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WYzth_zkevmL1 *WYzth_zkevmL1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.WYzth_zkevmL1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WYzth_zkevmL1 *WYzth_zkevmL1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.WYzth_zkevmL1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WYzth_zkevmL1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) AddressManager() (common.Address, error) {
	return _WYzth_zkevmL1.Contract.AddressManager(&_WYzth_zkevmL1.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) AddressManager() (common.Address, error) {
	return _WYzth_zkevmL1.Contract.AddressManager(&_WYzth_zkevmL1.CallOpts)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetBlock(opts *bind.CallOpts, blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getBlock", blockId)

	outstruct := new(struct {
		MetaHash   [32]byte
		Proposer   common.Address
		ProposedAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MetaHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Proposer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ProposedAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetBlock(blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	return _WYzth_zkevmL1.Contract.GetBlock(&_WYzth_zkevmL1.CallOpts, blockId)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetBlock(blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	return _WYzth_zkevmL1.Contract.GetBlock(&_WYzth_zkevmL1.CallOpts, blockId)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetBlockFee(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getBlockFee")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetBlockFee() (uint64, error) {
	return _WYzth_zkevmL1.Contract.GetBlockFee(&_WYzth_zkevmL1.CallOpts)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetBlockFee() (uint64, error) {
	return _WYzth_zkevmL1.Contract.GetBlockFee(&_WYzth_zkevmL1.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint64,uint8,bool))
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetConfig(opts *bind.CallOpts) (WYzth_zkevmDataConfig, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(WYzth_zkevmDataConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(WYzth_zkevmDataConfig)).(*WYzth_zkevmDataConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint64,uint8,bool))
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetConfig() (WYzth_zkevmDataConfig, error) {
	return _WYzth_zkevmL1.Contract.GetConfig(&_WYzth_zkevmL1.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint64,uint8,bool))
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetConfig() (WYzth_zkevmDataConfig, error) {
	return _WYzth_zkevmL1.Contract.GetConfig(&_WYzth_zkevmL1.CallOpts)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetCrossChainBlockHash(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getCrossChainBlockHash", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetCrossChainBlockHash(blockId *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL1.Contract.GetCrossChainBlockHash(&_WYzth_zkevmL1.CallOpts, blockId)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetCrossChainBlockHash(blockId *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL1.Contract.GetCrossChainBlockHash(&_WYzth_zkevmL1.CallOpts, blockId)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetCrossChainSignalRoot(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getCrossChainSignalRoot", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetCrossChainSignalRoot(blockId *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL1.Contract.GetCrossChainSignalRoot(&_WYzth_zkevmL1.CallOpts, blockId)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetCrossChainSignalRoot(blockId *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL1.Contract.GetCrossChainSignalRoot(&_WYzth_zkevmL1.CallOpts, blockId)
}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetForkChoice(opts *bind.CallOpts, blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (WYzth_zkevmDataForkChoice, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getForkChoice", blockId, parentHash, parentGasUsed)

	if err != nil {
		return *new(WYzth_zkevmDataForkChoice), err
	}

	out0 := *abi.ConvertType(out[0], new(WYzth_zkevmDataForkChoice)).(*WYzth_zkevmDataForkChoice)

	return out0, err

}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetForkChoice(blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (WYzth_zkevmDataForkChoice, error) {
	return _WYzth_zkevmL1.Contract.GetForkChoice(&_WYzth_zkevmL1.CallOpts, blockId, parentHash, parentGasUsed)
}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetForkChoice(blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (WYzth_zkevmDataForkChoice, error) {
	return _WYzth_zkevmL1.Contract.GetForkChoice(&_WYzth_zkevmL1.CallOpts, blockId, parentHash, parentGasUsed)
}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint64)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetProofReward(opts *bind.CallOpts, provenAt uint64, proposedAt uint64) (uint64, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getProofReward", provenAt, proposedAt)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint64)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetProofReward(provenAt uint64, proposedAt uint64) (uint64, error) {
	return _WYzth_zkevmL1.Contract.GetProofReward(&_WYzth_zkevmL1.CallOpts, provenAt, proposedAt)
}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint64)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetProofReward(provenAt uint64, proposedAt uint64) (uint64, error) {
	return _WYzth_zkevmL1.Contract.GetProofReward(&_WYzth_zkevmL1.CallOpts, provenAt, proposedAt)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetStateVariables(opts *bind.CallOpts) (WYzth_zkevmDataStateVariables, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getStateVariables")

	if err != nil {
		return *new(WYzth_zkevmDataStateVariables), err
	}

	out0 := *abi.ConvertType(out[0], new(WYzth_zkevmDataStateVariables)).(*WYzth_zkevmDataStateVariables)

	return out0, err

}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetStateVariables() (WYzth_zkevmDataStateVariables, error) {
	return _WYzth_zkevmL1.Contract.GetStateVariables(&_WYzth_zkevmL1.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetStateVariables() (WYzth_zkevmDataStateVariables, error) {
	return _WYzth_zkevmL1.Contract.GetStateVariables(&_WYzth_zkevmL1.CallOpts)
}

// GetWYzth_zkevmTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getWYzth_zkevmTokenBalance(address addr) view returns(uint256)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetWYzth_zkevmTokenBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getWYzth_zkevmTokenBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWYzth_zkevmTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getWYzth_zkevmTokenBalance(address addr) view returns(uint256)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetWYzth_zkevmTokenBalance(addr common.Address) (*big.Int, error) {
	return _WYzth_zkevmL1.Contract.GetWYzth_zkevmTokenBalance(&_WYzth_zkevmL1.CallOpts, addr)
}

// GetWYzth_zkevmTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getWYzth_zkevmTokenBalance(address addr) view returns(uint256)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetWYzth_zkevmTokenBalance(addr common.Address) (*big.Int, error) {
	return _WYzth_zkevmL1.Contract.GetWYzth_zkevmTokenBalance(&_WYzth_zkevmL1.CallOpts, addr)
}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) GetVerifierName(opts *bind.CallOpts, id uint16) ([32]byte, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "getVerifierName", id)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) GetVerifierName(id uint16) ([32]byte, error) {
	return _WYzth_zkevmL1.Contract.GetVerifierName(&_WYzth_zkevmL1.CallOpts, id)
}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) GetVerifierName(id uint16) ([32]byte, error) {
	return _WYzth_zkevmL1.Contract.GetVerifierName(&_WYzth_zkevmL1.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) Owner() (common.Address, error) {
	return _WYzth_zkevmL1.Contract.Owner(&_WYzth_zkevmL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) Owner() (common.Address, error) {
	return _WYzth_zkevmL1.Contract.Owner(&_WYzth_zkevmL1.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _WYzth_zkevmL1.Contract.Resolve(&_WYzth_zkevmL1.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _WYzth_zkevmL1.Contract.Resolve(&_WYzth_zkevmL1.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _WYzth_zkevmL1.Contract.Resolve0(&_WYzth_zkevmL1.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _WYzth_zkevmL1.Contract.Resolve0(&_WYzth_zkevmL1.CallOpts, name, allowZeroAddress)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved71, uint64 __reserved72, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 __reserved91)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Caller) State(opts *bind.CallOpts) (struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved71              uint64
	Reserved72              uint64
	AccProposedAt           uint64
	AccBlockFees            uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	BlockFee                uint64
	ProofTimeIssued         uint64
	LastVerifiedBlockId     uint64
	Reserved91              uint64
}, error) {
	var out []interface{}
	err := _WYzth_zkevmL1.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		GenesisHeight           uint64
		GenesisTimestamp        uint64
		Reserved71              uint64
		Reserved72              uint64
		AccProposedAt           uint64
		AccBlockFees            uint64
		NumBlocks               uint64
		NextEthDepositToProcess uint64
		BlockFee                uint64
		ProofTimeIssued         uint64
		LastVerifiedBlockId     uint64
		Reserved91              uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GenesisHeight = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.GenesisTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Reserved71 = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Reserved72 = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.AccProposedAt = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.AccBlockFees = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.NumBlocks = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.NextEthDepositToProcess = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.BlockFee = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.ProofTimeIssued = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.LastVerifiedBlockId = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.Reserved91 = *abi.ConvertType(out[11], new(uint64)).(*uint64)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved71, uint64 __reserved72, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 __reserved91)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) State() (struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved71              uint64
	Reserved72              uint64
	AccProposedAt           uint64
	AccBlockFees            uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	BlockFee                uint64
	ProofTimeIssued         uint64
	LastVerifiedBlockId     uint64
	Reserved91              uint64
}, error) {
	return _WYzth_zkevmL1.Contract.State(&_WYzth_zkevmL1.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved71, uint64 __reserved72, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 __reserved91)
func (_WYzth_zkevmL1 *WYzth_zkevmL1CallerSession) State() (struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved71              uint64
	Reserved72              uint64
	AccProposedAt           uint64
	AccBlockFees            uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	BlockFee                uint64
	ProofTimeIssued         uint64
	LastVerifiedBlockId     uint64
	Reserved91              uint64
}, error) {
	return _WYzth_zkevmL1.Contract.State(&_WYzth_zkevmL1.CallOpts)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) DepositEtherToL2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "depositEtherToL2")
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) DepositEtherToL2() (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.DepositEtherToL2(&_WYzth_zkevmL1.TransactOpts)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) DepositEtherToL2() (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.DepositEtherToL2(&_WYzth_zkevmL1.TransactOpts)
}

// DepositWYzth_zkevmToken is a paid mutator transaction binding the contract method 0x98f39aba.
//
// Solidity: function depositWYzth_zkevmToken(uint256 amount) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) DepositWYzth_zkevmToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "depositWYzth_zkevmToken", amount)
}

// DepositWYzth_zkevmToken is a paid mutator transaction binding the contract method 0x98f39aba.
//
// Solidity: function depositWYzth_zkevmToken(uint256 amount) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) DepositWYzth_zkevmToken(amount *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.DepositWYzth_zkevmToken(&_WYzth_zkevmL1.TransactOpts, amount)
}

// DepositWYzth_zkevmToken is a paid mutator transaction binding the contract method 0x98f39aba.
//
// Solidity: function depositWYzth_zkevmToken(uint256 amount) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) DepositWYzth_zkevmToken(amount *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.DepositWYzth_zkevmToken(&_WYzth_zkevmL1.TransactOpts, amount)
}

// Init is a paid mutator transaction binding the contract method 0x578c65a4.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeIssued) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeIssued uint64) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "init", _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeIssued)
}

// Init is a paid mutator transaction binding the contract method 0x578c65a4.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeIssued) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeIssued uint64) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.Init(&_WYzth_zkevmL1.TransactOpts, _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeIssued)
}

// Init is a paid mutator transaction binding the contract method 0x578c65a4.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeIssued) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeIssued uint64) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.Init(&_WYzth_zkevmL1.TransactOpts, _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeIssued)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) ProposeBlock(opts *bind.TransactOpts, input []byte, txList []byte) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "proposeBlock", input, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) ProposeBlock(input []byte, txList []byte) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.ProposeBlock(&_WYzth_zkevmL1.TransactOpts, input, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) ProposeBlock(input []byte, txList []byte) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.ProposeBlock(&_WYzth_zkevmL1.TransactOpts, input, txList)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) ProveBlock(opts *bind.TransactOpts, blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "proveBlock", blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) ProveBlock(blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.ProveBlock(&_WYzth_zkevmL1.TransactOpts, blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) ProveBlock(blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.ProveBlock(&_WYzth_zkevmL1.TransactOpts, blockId, input)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) RenounceOwnership() (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.RenounceOwnership(&_WYzth_zkevmL1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.RenounceOwnership(&_WYzth_zkevmL1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.TransferOwnership(&_WYzth_zkevmL1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.TransferOwnership(&_WYzth_zkevmL1.TransactOpts, newOwner)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) VerifyBlocks(opts *bind.TransactOpts, maxBlocks *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "verifyBlocks", maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) VerifyBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.VerifyBlocks(&_WYzth_zkevmL1.TransactOpts, maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) VerifyBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.VerifyBlocks(&_WYzth_zkevmL1.TransactOpts, maxBlocks)
}

// WithdrawWYzth_zkevmToken is a paid mutator transaction binding the contract method 0x5043f059.
//
// Solidity: function withdrawWYzth_zkevmToken(uint256 amount) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) WithdrawWYzth_zkevmToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.Transact(opts, "withdrawWYzth_zkevmToken", amount)
}

// WithdrawWYzth_zkevmToken is a paid mutator transaction binding the contract method 0x5043f059.
//
// Solidity: function withdrawWYzth_zkevmToken(uint256 amount) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) WithdrawWYzth_zkevmToken(amount *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.WithdrawWYzth_zkevmToken(&_WYzth_zkevmL1.TransactOpts, amount)
}

// WithdrawWYzth_zkevmToken is a paid mutator transaction binding the contract method 0x5043f059.
//
// Solidity: function withdrawWYzth_zkevmToken(uint256 amount) returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) WithdrawWYzth_zkevmToken(amount *big.Int) (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.WithdrawWYzth_zkevmToken(&_WYzth_zkevmL1.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WYzth_zkevmL1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1Session) Receive() (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.Receive(&_WYzth_zkevmL1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WYzth_zkevmL1 *WYzth_zkevmL1TransactorSession) Receive() (*types.Transaction, error) {
	return _WYzth_zkevmL1.Contract.Receive(&_WYzth_zkevmL1.TransactOpts)
}

// WYzth_zkevmL1BlockProposedIterator is returned from FilterBlockProposed and is used to iterate over the raw logs and unpacked data for BlockProposed events raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1BlockProposedIterator struct {
	Event *WYzth_zkevmL1BlockProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WYzth_zkevmL1BlockProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL1BlockProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WYzth_zkevmL1BlockProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WYzth_zkevmL1BlockProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL1BlockProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL1BlockProposed represents a BlockProposed event raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1BlockProposed struct {
	Id   *big.Int
	Meta WYzth_zkevmDataBlockMetadata
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x31cc3eebb5aea55796fe5ba252fe1833fd17d063cb6e83634758fa58e7181535.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) FilterBlockProposed(opts *bind.FilterOpts, id []*big.Int) (*WYzth_zkevmL1BlockProposedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.FilterLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1BlockProposedIterator{contract: _WYzth_zkevmL1.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0x31cc3eebb5aea55796fe5ba252fe1833fd17d063cb6e83634758fa58e7181535.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL1BlockProposed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.WatchLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL1BlockProposed)
				if err := _WYzth_zkevmL1.contract.UnpackLog(event, "BlockProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockProposed is a log parse operation binding the contract event 0x31cc3eebb5aea55796fe5ba252fe1833fd17d063cb6e83634758fa58e7181535.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) ParseBlockProposed(log types.Log) (*WYzth_zkevmL1BlockProposed, error) {
	event := new(WYzth_zkevmL1BlockProposed)
	if err := _WYzth_zkevmL1.contract.UnpackLog(event, "BlockProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL1BlockProvenIterator is returned from FilterBlockProven and is used to iterate over the raw logs and unpacked data for BlockProven events raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1BlockProvenIterator struct {
	Event *WYzth_zkevmL1BlockProven // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WYzth_zkevmL1BlockProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL1BlockProven)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WYzth_zkevmL1BlockProven)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WYzth_zkevmL1BlockProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL1BlockProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL1BlockProven represents a BlockProven event raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1BlockProven struct {
	Id            *big.Int
	ParentHash    [32]byte
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Prover        common.Address
	ParentGasUsed uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockProven is a free log retrieval operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) FilterBlockProven(opts *bind.FilterOpts, id []*big.Int) (*WYzth_zkevmL1BlockProvenIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.FilterLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1BlockProvenIterator{contract: _WYzth_zkevmL1.contract, event: "BlockProven", logs: logs, sub: sub}, nil
}

// WatchBlockProven is a free log subscription operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) WatchBlockProven(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL1BlockProven, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.WatchLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL1BlockProven)
				if err := _WYzth_zkevmL1.contract.UnpackLog(event, "BlockProven", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockProven is a log parse operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) ParseBlockProven(log types.Log) (*WYzth_zkevmL1BlockProven, error) {
	event := new(WYzth_zkevmL1BlockProven)
	if err := _WYzth_zkevmL1.contract.UnpackLog(event, "BlockProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL1BlockVerifiedIterator is returned from FilterBlockVerified and is used to iterate over the raw logs and unpacked data for BlockVerified events raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1BlockVerifiedIterator struct {
	Event *WYzth_zkevmL1BlockVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WYzth_zkevmL1BlockVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL1BlockVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WYzth_zkevmL1BlockVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WYzth_zkevmL1BlockVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL1BlockVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL1BlockVerified represents a BlockVerified event raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1BlockVerified struct {
	Id        *big.Int
	BlockHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified is a free log retrieval operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) FilterBlockVerified(opts *bind.FilterOpts, id []*big.Int) (*WYzth_zkevmL1BlockVerifiedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.FilterLogs(opts, "BlockVerified", idRule)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1BlockVerifiedIterator{contract: _WYzth_zkevmL1.contract, event: "BlockVerified", logs: logs, sub: sub}, nil
}

// WatchBlockVerified is a free log subscription operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) WatchBlockVerified(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL1BlockVerified, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.WatchLogs(opts, "BlockVerified", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL1BlockVerified)
				if err := _WYzth_zkevmL1.contract.UnpackLog(event, "BlockVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockVerified is a log parse operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) ParseBlockVerified(log types.Log) (*WYzth_zkevmL1BlockVerified, error) {
	event := new(WYzth_zkevmL1BlockVerified)
	if err := _WYzth_zkevmL1.contract.UnpackLog(event, "BlockVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL1CrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1CrossChainSyncedIterator struct {
	Event *WYzth_zkevmL1CrossChainSynced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WYzth_zkevmL1CrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL1CrossChainSynced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WYzth_zkevmL1CrossChainSynced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WYzth_zkevmL1CrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL1CrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL1CrossChainSynced represents a CrossChainSynced event raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1CrossChainSynced struct {
	SrcHeight  *big.Int
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) FilterCrossChainSynced(opts *bind.FilterOpts, srcHeight []*big.Int) (*WYzth_zkevmL1CrossChainSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.FilterLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1CrossChainSyncedIterator{contract: _WYzth_zkevmL1.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL1CrossChainSynced, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.WatchLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL1CrossChainSynced)
				if err := _WYzth_zkevmL1.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainSynced is a log parse operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) ParseCrossChainSynced(log types.Log) (*WYzth_zkevmL1CrossChainSynced, error) {
	event := new(WYzth_zkevmL1CrossChainSynced)
	if err := _WYzth_zkevmL1.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL1EthDepositedIterator is returned from FilterEthDeposited and is used to iterate over the raw logs and unpacked data for EthDeposited events raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1EthDepositedIterator struct {
	Event *WYzth_zkevmL1EthDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WYzth_zkevmL1EthDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL1EthDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WYzth_zkevmL1EthDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WYzth_zkevmL1EthDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL1EthDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL1EthDeposited represents a EthDeposited event raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1EthDeposited struct {
	Deposit WYzth_zkevmDataEthDeposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEthDeposited is a free log retrieval operation binding the contract event 0x1c146ddfc652d3b24f1c37ed1cabbb690bf0e198aea624f49211500467182eba.
//
// Solidity: event EthDeposited((address,uint96) deposit)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) FilterEthDeposited(opts *bind.FilterOpts) (*WYzth_zkevmL1EthDepositedIterator, error) {

	logs, sub, err := _WYzth_zkevmL1.contract.FilterLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1EthDepositedIterator{contract: _WYzth_zkevmL1.contract, event: "EthDeposited", logs: logs, sub: sub}, nil
}

// WatchEthDeposited is a free log subscription operation binding the contract event 0x1c146ddfc652d3b24f1c37ed1cabbb690bf0e198aea624f49211500467182eba.
//
// Solidity: event EthDeposited((address,uint96) deposit)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) WatchEthDeposited(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL1EthDeposited) (event.Subscription, error) {

	logs, sub, err := _WYzth_zkevmL1.contract.WatchLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL1EthDeposited)
				if err := _WYzth_zkevmL1.contract.UnpackLog(event, "EthDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthDeposited is a log parse operation binding the contract event 0x1c146ddfc652d3b24f1c37ed1cabbb690bf0e198aea624f49211500467182eba.
//
// Solidity: event EthDeposited((address,uint96) deposit)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) ParseEthDeposited(log types.Log) (*WYzth_zkevmL1EthDeposited, error) {
	event := new(WYzth_zkevmL1EthDeposited)
	if err := _WYzth_zkevmL1.contract.UnpackLog(event, "EthDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1InitializedIterator struct {
	Event *WYzth_zkevmL1Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WYzth_zkevmL1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL1Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WYzth_zkevmL1Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WYzth_zkevmL1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL1Initialized represents a Initialized event raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) FilterInitialized(opts *bind.FilterOpts) (*WYzth_zkevmL1InitializedIterator, error) {

	logs, sub, err := _WYzth_zkevmL1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1InitializedIterator{contract: _WYzth_zkevmL1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL1Initialized) (event.Subscription, error) {

	logs, sub, err := _WYzth_zkevmL1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL1Initialized)
				if err := _WYzth_zkevmL1.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) ParseInitialized(log types.Log) (*WYzth_zkevmL1Initialized, error) {
	event := new(WYzth_zkevmL1Initialized)
	if err := _WYzth_zkevmL1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1OwnershipTransferredIterator struct {
	Event *WYzth_zkevmL1OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WYzth_zkevmL1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL1OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WYzth_zkevmL1OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WYzth_zkevmL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL1OwnershipTransferred represents a OwnershipTransferred event raised by the WYzth_zkevmL1 contract.
type WYzth_zkevmL1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WYzth_zkevmL1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL1OwnershipTransferredIterator{contract: _WYzth_zkevmL1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WYzth_zkevmL1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL1OwnershipTransferred)
				if err := _WYzth_zkevmL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WYzth_zkevmL1 *WYzth_zkevmL1Filterer) ParseOwnershipTransferred(log types.Log) (*WYzth_zkevmL1OwnershipTransferred, error) {
	event := new(WYzth_zkevmL1OwnershipTransferred)
	if err := _WYzth_zkevmL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
