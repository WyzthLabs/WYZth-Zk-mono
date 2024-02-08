// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.20;

import { AddressResolver } from "../common/AddressResolver.sol";
import { EssentialContract } from "../common/EssentialContract.sol";
import { ICrossChainSync } from "../common/ICrossChainSync.sol";
import { Proxied } from "../common/Proxied.sol";
import { LibEthDepositing } from "./libs/LibEthDepositing.sol";
import { LibProposing } from "./libs/LibProposing.sol";
import { LibProving } from "./libs/LibProving.sol";
import { LibUtils } from "./libs/LibUtils.sol";
import { LibVerifying } from "./libs/LibVerifying.sol";
import { WYzth_zkevmConfig } from "./WYzth_zkevmConfig.sol";
import { WYzth_zkevmErrors } from "./WYzth_zkevmErrors.sol";
import { WYzth_zkevmData } from "./WYzth_zkevmData.sol";
import { WYzth_zkevmEvents } from "./WYzth_zkevmEvents.sol";

/// @custom:security-contact hello@wyzth_zkevm.xyz
contract WYzth_zkevmL1 is
    EssentialContract,
    ICrossChainSync,
    WYzth_zkevmEvents,
    WYzth_zkevmErrors
{
    using LibUtils for WYzth_zkevmData.State;

    WYzth_zkevmData.State public state;
    uint256[100] private __gap;

    receive() external payable {
        depositEtherToL2(address(0));
    }

    /**
     * Initialize the rollup.
     *
     * @param _addressManager The AddressManager address.
     * @param _genesisBlockHash The block hash of the genesis block.
     * @param _initBlockFee Initial (reasonable) block fee value.
     */
    function init(
        address _addressManager,
        bytes32 _genesisBlockHash,
        uint64 _initBlockFee
    )
        external
        initializer
    {
        EssentialContract._init(_addressManager);
        LibVerifying.init({
            state: state,
            config: getConfig(),
            genesisBlockHash: _genesisBlockHash,
            initBlockFee: _initBlockFee
        });
    }

    /**
     * Propose a WYzth_zkevm L2 block.
     *
     * @param input An abi-encoded BlockMetadataInput that the actual L2
     *        block header must satisfy.
     * @param txList A list of transactions in this block, encoded with RLP.
     *        Note, in the corresponding L2 block an _anchor transaction_
     *        will be the first transaction in the block -- if there are
     *        `n` transactions in `txList`, then there will be up to `n + 1`
     *        transactions in the L2 block.
     */
    function proposeBlock(
        bytes calldata input,
        bytes calldata txList
    )
        external
        nonReentrant
        returns (WYzth_zkevmData.BlockMetadata memory meta)
    {
        WYzth_zkevmData.Config memory config = getConfig();
        meta = LibProposing.proposeBlock({
            state: state,
            config: config,
            resolver: AddressResolver(this),
            input: abi.decode(input, (WYzth_zkevmData.BlockMetadataInput)),
            txList: txList
        });
        if (config.maxVerificationsPerTx > 0) {
            LibVerifying.verifyBlocks({
                state: state,
                config: config,
                resolver: AddressResolver(this),
                maxBlocks: config.maxVerificationsPerTx
            });
        }
    }

    /**
     * Prove a block with a zero-knowledge proof.
     *
     * @param blockId The index of the block to prove. This is also used
     *        to select the right implementation version.
     * @param input An abi-encoded WYzth_zkevmData.BlockEvidence object.
     */
    function proveBlock(
        uint256 blockId,
        bytes calldata input
    )
        external
        nonReentrant
    {
        WYzth_zkevmData.Config memory config = getConfig();
        LibProving.proveBlock({
            state: state,
            config: config,
            resolver: AddressResolver(this),
            blockId: blockId,
            evidence: abi.decode(input, (WYzth_zkevmData.BlockEvidence))
        });
        if (config.maxVerificationsPerTx > 0) {
            LibVerifying.verifyBlocks({
                state: state,
                config: config,
                resolver: AddressResolver(this),
                maxBlocks: config.maxVerificationsPerTx
            });
        }
    }

    /**
     * Verify up to N blocks.
     * @param maxBlocks Max number of blocks to verify.
     */
    function verifyBlocks(uint256 maxBlocks) external nonReentrant {
        if (maxBlocks == 0) revert L1_INVALID_PARAM();
        LibVerifying.verifyBlocks({
            state: state,
            config: getConfig(),
            resolver: AddressResolver(this),
            maxBlocks: maxBlocks
        });
    }

    function depositEtherToL2(address recipient) public payable {
        LibEthDepositing.depositEtherToL2({
            state: state,
            config: getConfig(),
            resolver: AddressResolver(this),
            recipient: recipient
        });
    }

    function canDepositEthToL2(uint256 amount) public view returns (bool) {
        return LibEthDepositing.canDepositEthToL2({
            state: state,
            config: getConfig(),
            amount: amount
        });
    }

    function getWYzth_zkevmTokenBalance(address addr) public view returns (uint256) {
        return state.wyzth_zkevmTokenBalances[addr];
    }

    function getBlockFee() public view returns (uint64) {
        return state.blockFee;
    }

    function getBlock(uint256 blockId)
        public
        view
        returns (bytes32 _metaHash, address _proposer, uint64 _proposedAt)
    {
        WYzth_zkevmData.Block storage blk = LibProposing.getBlock({
            state: state,
            config: getConfig(),
            blockId: blockId
        });
        _metaHash = blk.metaHash;
        _proposer = blk.proposer;
        _proposedAt = blk.proposedAt;
    }

    function getForkChoice(
        uint256 blockId,
        bytes32 parentHash,
        uint32 parentGasUsed
    )
        public
        view
        returns (WYzth_zkevmData.ForkChoice memory)
    {
        return LibProving.getForkChoice({
            state: state,
            config: getConfig(),
            blockId: blockId,
            parentHash: parentHash,
            parentGasUsed: parentGasUsed
        });
    }

    function getCrossChainBlockHash(uint256 blockId)
        public
        view
        override
        returns (bytes32)
    {
        (bool found, WYzth_zkevmData.Block storage blk) = LibUtils.getL2ChainData({
            state: state,
            config: getConfig(),
            blockId: blockId
        });
        return found
            ? blk.forkChoices[blk.verifiedForkChoiceId].blockHash
            : bytes32(0);
    }

    function getCrossChainSignalRoot(uint256 blockId)
        public
        view
        override
        returns (bytes32)
    {
        (bool found, WYzth_zkevmData.Block storage blk) = LibUtils.getL2ChainData({
            state: state,
            config: getConfig(),
            blockId: blockId
        });

        return found
            ? blk.forkChoices[blk.verifiedForkChoiceId].signalRoot
            : bytes32(0);
    }

    function getStateVariables()
        public
        view
        returns (WYzth_zkevmData.StateVariables memory)
    {
        return state.getStateVariables();
    }

    function getConfig()
        public
        pure
        virtual
        returns (WYzth_zkevmData.Config memory)
    {
        return WYzth_zkevmConfig.getConfig();
    }

    function getVerifierName(uint16 id) public pure returns (bytes32) {
        return LibUtils.getVerifierName(id);
    }
}

contract ProxiedWYzth_zkevmL1 is Proxied, WYzth_zkevmL1 { }
