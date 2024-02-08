import type { ethers } from "ethers";
import WYzth_zkevmL1 from "../constants/abi/WYzth_zkevmL1";
import { getStateVariables } from "./getStateVariables";

export const getPendingBlocks = async (
  provider: ethers.providers.JsonRpcProvider,
  contractAddress: string
): Promise<number> => {
  const stateVariables = await getStateVariables(provider, contractAddress);
  const nextBlockId = stateVariables.numBlocks;
  const lastBlockId = stateVariables.lastVerifiedBlockId;
  return nextBlockId - lastBlockId - 1;
};
