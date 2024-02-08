// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "forge-std/console2.sol";
import
    "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeCastUpgradeable.sol";
import "../../contracts/L1/WYzth_zkevmL1.sol";
import "./UpgradeScript.s.sol";

contract UpgradeWYzth_zkevmL1 is UpgradeScript {
    function run() external setUp {
        WYzth_zkevmL1 newWYzth_zkevmL1 = new ProxiedWYzth_zkevmL1();
        proxy.upgradeTo(address(newWYzth_zkevmL1));
        console2.log(
            "proxy upgraded wyzth_zkevm L1 implementation to", address(newWYzth_zkevmL1)
        );
    }
}
