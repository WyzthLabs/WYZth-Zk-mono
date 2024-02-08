import { Contract, ethers } from "ethers";
import WYzth_zkevmL1 from "../constants/abi/WYzth_zkevmL1";

export const getConfig = async (
  provider: ethers.providers.JsonRpcProvider,
  contractAddress: string
) => {
  const contract: Contract = new Contract(contractAddress, WYzth_zkevmL1, provider);
  const config = await contract.getConfig();
  return config;
};
