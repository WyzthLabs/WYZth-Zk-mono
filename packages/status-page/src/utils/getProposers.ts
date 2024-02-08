import { Contract, ethers } from "ethers";
import WYzth_zkevmL1 from "../constants/abi/WYzth_zkevmL1";

export const getProposers = async (
  provider: ethers.providers.JsonRpcProvider,
  contractAddress: string
): Promise<number> => {
  const contract: Contract = new Contract(contractAddress, WYzth_zkevmL1, provider);
  let events = [];
  const latestBlock = await provider.getBlockNumber();
  const batchSize = 1000;
  for (let i = 0; i < latestBlock; i += batchSize) {
    const end = i + batchSize > latestBlock ? latestBlock : i + batchSize;
    const e = await contract.queryFilter("BlockProposed", i, end);
    events = events.concat(e);
  }

  const proposers = [];
  events.map((event) => {
    if (!proposers.includes(event.args.meta.beneficiary)) {
      proposers.push(event.args.meta.beneficiary);
    }
  });

  return proposers.length;
};
