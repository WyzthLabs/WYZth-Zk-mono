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
import "../contracts/L1/WYzth_zkevmToken.sol";
import "../contracts/L1/WYzth_zkevmL1.sol";
import "../contracts/bridge/Bridge.sol";
import "../contracts/bridge/TokenVault.sol";
import "../contracts/signal/SignalService.sol";
import "../contracts/common/AddressManager.sol";
import "../contracts/test/erc20/FreeMintERC20.sol";
import "../contracts/test/erc20/MayFailFreeMintERC20.sol";

contract DeployOnL1 is Script {
    using SafeCastUpgradeable for uint256;

    bytes32 public genesisHash = vm.envBytes32("L2_GENESIS_HASH");

    uint256 public deployerPrivateKey = vm.envUint("PRIVATE_KEY");

    address public wyzth_zkevmL2Address = vm.envAddress("WYZTH_ZKEVM_L2_ADDRESS");

    address public l2SignalService = vm.envAddress("L2_SIGNAL_SERVICE");

    address public owner = vm.envAddress("OWNER");

    address public oracleProver = vm.envAddress("ORACLE_PROVER");
    address public systemProver = vm.envAddress("SYSTEM_PROVER");

    address public sharedSignalService = vm.envAddress("SHARED_SIGNAL_SERVICE");

    address public treasury = vm.envAddress("TREASURY");

    address public wyzth_zkevmTokenPremintRecipient =
        vm.envAddress("WYZTH_ZKEVM_TOKEN_PREMINT_RECIPIENT");

    uint256 public wyzth_zkevmTokenPremintAmount =
        vm.envUint("WYZTH_ZKEVM_TOKEN_PREMINT_AMOUNT");

    WYzth_zkevmL1 wyzth_zkevmL1;
    address public addressManagerProxy;

    error FAILED_TO_DEPLOY_PLONK_VERIFIER(string contractPath);

    function run() external {
        require(owner != address(0), "owner is zero");
        require(wyzth_zkevmL2Address != address(0), "wyzth_zkevmL2Address is zero");
        require(l2SignalService != address(0), "l2SignalService is zero");
        require(treasury != address(0), "treasury is zero");
        require(
            wyzth_zkevmTokenPremintRecipient != address(0),
            "wyzth_zkevmTokenPremintRecipient is zero"
        );
        require(wyzth_zkevmTokenPremintAmount < type(uint64).max, "premint too large");

        vm.startBroadcast(deployerPrivateKey);

        // AddressManager
        AddressManager addressManager = new ProxiedAddressManager();
        addressManagerProxy = deployProxy(
            "address_manager",
            address(addressManager),
            bytes.concat(addressManager.init.selector)
        );

        // WYzth_zkevmL1
        wyzth_zkevmL1 = new ProxiedWYzth_zkevmL1();
        uint256 l2ChainId = wyzth_zkevmL1.getConfig().chainId;
        require(l2ChainId != block.chainid, "same chainid");

        setAddress(l2ChainId, "wyzth_zkevm", wyzth_zkevmL2Address);
        setAddress(l2ChainId, "signal_service", l2SignalService);
        setAddress("oracle_prover", oracleProver);
        setAddress("system_prover", systemProver);
        setAddress(l2ChainId, "treasury", treasury);

        // WYzth_zkevmToken
        WYzth_zkevmToken wyzth_zkevmToken = new ProxiedWYzth_zkevmToken();

        address[] memory premintRecipients = new address[](1);
        uint256[] memory premintAmounts = new uint256[](1);
        premintRecipients[0] = wyzth_zkevmTokenPremintRecipient;
        premintAmounts[0] = wyzth_zkevmTokenPremintAmount;

        deployProxy(
            "wyzth_zkevm_token",
            address(wyzth_zkevmToken),
            bytes.concat(
                wyzth_zkevmToken.init.selector,
                abi.encode(
                    addressManagerProxy,
                    "WYzth_zkevm Token",
                    "TKO",
                    premintRecipients,
                    premintAmounts
                )
            )
        );

        // HorseToken && BullToken
        address horseToken = address(new FreeMintERC20("Horse Token", "HORSE"));
        console2.log("HorseToken", horseToken);

        address bullToken =
            address(new MayFailFreeMintERC20("Bull Token", "BLL"));
        console2.log("BullToken", bullToken);

        uint64 feeBase = uint64(1) ** wyzth_zkevmToken.decimals();

        address wyzth_zkevmL1Proxy = deployProxy(
            "wyzth_zkevm",
            address(wyzth_zkevmL1),
            bytes.concat(
                wyzth_zkevmL1.init.selector,
                abi.encode(addressManagerProxy, genesisHash, feeBase)
            )
        );
        setAddress("wyzth_zkevm", wyzth_zkevmL1Proxy);
        setAddress("proto_broker", wyzth_zkevmL1Proxy);

        // Bridge
        Bridge bridge = new ProxiedBridge();
        deployProxy(
            "bridge",
            address(bridge),
            bytes.concat(bridge.init.selector, abi.encode(addressManagerProxy))
        );

        // TokenVault
        TokenVault tokenVault = new ProxiedTokenVault();
        deployProxy(
            "token_vault",
            address(tokenVault),
            bytes.concat(
                tokenVault.init.selector, abi.encode(addressManagerProxy)
            )
        );

        // SignalService
        if (sharedSignalService == address(0)) {
            SignalService signalService = new ProxiedSignalService();
            deployProxy(
                "signal_service",
                address(signalService),
                bytes.concat(
                    signalService.init.selector, abi.encode(addressManagerProxy)
                )
            );
        } else {
            console2.log(
                "Warining: using shared signal service: ", sharedSignalService
            );
            setAddress("signal_service", sharedSignalService);
        }

        // PlonkVerifier
        deployPlonkVerifiers();

        vm.stopBroadcast();
    }

    function deployPlonkVerifiers() private {
        address[] memory plonkVerifiers = new address[](1);
        plonkVerifiers[0] =
            deployYulContract("contracts/libs/yul/PlonkVerifier.yulp");

        for (uint16 i = 0; i < plonkVerifiers.length; ++i) {
            setAddress(wyzth_zkevmL1.getVerifierName(i), plonkVerifiers[i]);
        }
    }

    function deployYulContract(string memory contractPath)
        private
        returns (address)
    {
        string[] memory cmds = new string[](3);
        cmds[0] = "bash";
        cmds[1] = "-c";
        cmds[2] = string.concat(
            vm.projectRoot(),
            "/bin/solc --yul --bin ",
            string.concat(vm.projectRoot(), "/", contractPath),
            " | grep -A1 Binary | tail -1"
        );

        bytes memory bytecode = vm.ffi(cmds);

        address deployedAddress;
        assembly {
            deployedAddress := create(0, add(bytecode, 0x20), mload(bytecode))
        }

        if (deployedAddress == address(0)) {
            revert FAILED_TO_DEPLOY_PLONK_VERIFIER(contractPath);
        }

        console2.log(contractPath, deployedAddress);

        return deployedAddress;
    }

    function deployProxy(
        string memory name,
        address implementation,
        bytes memory data
    )
        private
        returns (address proxy)
    {
        proxy = address(
            new TransparentUpgradeableProxy(implementation, owner, data)
        );

        console2.log(name, "(impl) ->", implementation);
        console2.log(name, "(proxy) ->", proxy);

        if (addressManagerProxy != address(0)) {
            setAddress(block.chainid, bytes32(bytes(name)), proxy);
        }

        vm.writeJson(
            vm.serializeAddress("deployment", name, proxy),
            string.concat(vm.projectRoot(), "/deployments/deploy_l1.json")
        );
    }

    function setAddress(bytes32 name, address addr) private {
        setAddress(block.chainid, name, addr);
    }

    function setAddress(uint256 chainId, bytes32 name, address addr) private {
        console2.log(chainId, uint256(name), "--->", addr);
        if (addr != address(0)) {
            AddressManager(addressManagerProxy).setAddress(chainId, name, addr);
        }
    }
}
