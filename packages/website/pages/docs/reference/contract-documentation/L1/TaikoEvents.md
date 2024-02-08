---
title: WYzth_zkevmEvents
---

## WYzth_zkevmEvents

### BlockProposed

```solidity
event BlockProposed(uint256 id, struct WYzth_zkevmData.BlockMetadata meta, uint64 blockFee)
```

### BlockProven

```solidity
event BlockProven(uint256 id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
```

### BlockVerified

```solidity
event BlockVerified(uint256 id, bytes32 blockHash, uint64 reward)
```

### EthDeposited

```solidity
event EthDeposited(struct WYzth_zkevmData.EthDeposit deposit)
```

### ProofParamsChanged

```solidity
event ProofParamsChanged(uint64 proofTimeTarget, uint64 proofTimeIssued, uint64 blockFee, uint16 adjustmentQuotient)
```
