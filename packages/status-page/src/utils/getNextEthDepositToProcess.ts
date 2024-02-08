import { Contract, ethers } from "ethers";
import WYzth_zkevmL1 from "../constants/abi/WYzth_zkevmL1";

export const getNextEthDepositToProcess = async (
  provider: ethers.providers.JsonRpcProvider,
  contractAddress: string
): Promise<number> => {
  const contract: Contract = new Contract(contractAddress, WYzth_zkevmL1, provider);
  const stateVariables = await contract.getStateVariables();
  return stateVariables.nextEthDepositToProcess.toNumber();
};
