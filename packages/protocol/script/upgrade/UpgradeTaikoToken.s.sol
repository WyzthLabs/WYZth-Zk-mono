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
import "../../contracts/L1/WYzth_zkevmToken.sol";
import "./UpgradeScript.s.sol";

contract UpgradeWYzth_zkevmToken is UpgradeScript {
    function run() external setUp {
        WYzth_zkevmToken newWYzth_zkevmToken = new ProxiedWYzth_zkevmToken();
        proxy.upgradeTo(address(newWYzth_zkevmToken));
        console2.log(
            "proxy upgraded WYzth_zkevmToken implementation to",
            address(newWYzth_zkevmToken)
        );
    }
}
