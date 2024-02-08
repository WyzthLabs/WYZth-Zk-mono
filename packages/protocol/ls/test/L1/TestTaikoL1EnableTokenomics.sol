// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.20;

import { WYzth_zkevmL1 } from "../../L1/WYzth_zkevmL1.sol";
import { WYzth_zkevmData } from "../../L1/WYzth_zkevmData.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";

contract TestWYzth_zkevmL1EnableTokenomics is WYzth_zkevmL1 {
    function getConfig()
        public
        pure
        override
        returns (WYzth_zkevmData.Config memory config)
    {
        config.chainId = 167;
        // up to 2048 pending blocks
        config.maxNumProposedBlocks = 6;
        config.ringBufferSize = 8;
        // This number is calculated from maxNumProposedBlocks to make
        // the 'the maximum value of the multiplier' close to 20.0
        config.maxVerificationsPerTx = 0; // dont verify blocks automatically
        config.blockMaxGasLimit = 30_000_000;
        config.maxTransactionsPerBlock = 20;
        config.maxBytesPerTxList = 120_000;
    }
}
