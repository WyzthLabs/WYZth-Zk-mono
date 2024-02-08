import { Contract, ethers } from "ethers";
import WYzth_zkevmL1 from "../constants/abi/WYzth_zkevmL1";

export const getLatestSyncedHeader = async (
  provider: ethers.providers.JsonRpcProvider,
  contractAddress: string
): Promise<string> => {
  const contract: Contract = new Contract(contractAddress, WYzth_zkevmL1, provider);
  return await contract.getCrossChainBlockHash(0);
};
