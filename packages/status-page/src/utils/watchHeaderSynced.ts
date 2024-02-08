import { Contract, ethers } from "ethers";
import WYzth_zkevmL1 from "../constants/abi/WYzth_zkevmL1";

export const watchHeaderSynced = async (
  provider: ethers.providers.JsonRpcProvider,
  wyzth_zkevmL1Address: string,
  onEvent: (value: string | number | boolean) => void
) => {
  const contract: Contract = new Contract(wyzth_zkevmL1Address, WYzth_zkevmL1, provider);
  const listener = (lastVerifiedBlockId, blockHash, signalRoot) => {
    onEvent(blockHash);
  };
  contract.on("CrossChainSynced", listener);

  return () => contract.off("CrossChainSynced", listener);
};
