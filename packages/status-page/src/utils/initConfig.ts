import { ethers } from "ethers";
import { Layer } from "../domain/layer";

export function initConfig(layer: Layer) {
  const l1Provider = new ethers.providers.StaticJsonRpcProvider(
    layer === Layer.Two
      ? import.meta.env.VITE_L1_RPC_URL
      : import.meta.env.VITE_L2_RPC_URL
  );
  const l2Provider = new ethers.providers.StaticJsonRpcProvider(
    layer === Layer.Two
      ? import.meta.env.VITE_L2_RPC_URL
      : import.meta.env.VITE_L3_RPC_URL
  );

  const l1WYzth_zkevmAddress =
    layer === Layer.Two
      ? import.meta.env.VITE_L2_WYZTH_ZKEVM_L1_ADDRESS
      : import.meta.env.VITE_L3_WYZTH_ZKEVM_L1_ADDRESS;
  const l2WYzth_zkevmAddress =
    layer === Layer.Two
      ? import.meta.env.VITE_L2_WYZTH_ZKEVM_L2_ADDRESS
      : import.meta.env.VITE_L3_WYZTH_ZKEVM_L2_ADDRESS;
  const wyzth_zkevmTokenAddress = import.meta.env.VITE_WYZTH_ZKEVM_TOKEN_ADDRESS;
  const l1ExplorerUrl = import.meta.env.VITE_L1_EXPLORER_URL;
  const l2ExplorerUrl =
    layer === Layer.Two
      ? import.meta.env.VITE_L2_EXPLORER_URL
      : import.meta.env.VITE_L3_EXPLORER_URL;
  const feeTokenSymbol = import.meta.env.VITE_FEE_TOKEN_SYMBOL || "TKO";
  const oracleProverAddress =
    import.meta.env.ORACLE_PROVER_ADDRESS ||
    "0x0000000000000000000000000000000000000000";
  const systemProverAddress =
    import.meta.env.SYSTEM_PROVER_ADDRESS ||
    "0x0000000000000000000000000000000000000001";
  const eventIndexerApiUrl =
    layer === Layer.Two
      ? import.meta.env.VITE_L2_EVENT_INDEXER_API_URL
      : import.meta.env.VITE_L3_EVENT_INDEXER_API_URL;
  return {
    l1Provider,
    l2Provider,
    l1WYzth_zkevmAddress,
    l2WYzth_zkevmAddress,
    wyzth_zkevmTokenAddress,
    l1ExplorerUrl,
    l2ExplorerUrl,
    feeTokenSymbol,
    oracleProverAddress,
    eventIndexerApiUrl,
    systemProverAddress,
  };
}
