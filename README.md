# ReflexNet (MindCell Protocol)

A Cosmos SDK-based decentralized protocol for AI model sharding, validation, and monetization.

## Overview

ReflexNet (MindCell) enables decentralized storage and execution of AI models through:
- **Model Sharding**: Split ML models into encrypted fragments
- **Pay-per-Inference**: On-chain billing for model inference requests
- **zkML Validation**: Zero-knowledge proofs for inference correctness
- **Incentive Mechanism**: Staking and slashing for node operators
- **Versioned Registry**: Track and manage model versions

## Architecture

The protocol consists of multiple Cosmos SDK modules:

- **ModelRegistry**: Register and manage AI models with metadata
- **ShardAllocator**: Assign model shards to validator nodes
- **InferenceGateway**: Route inference requests and verify zkML proofs
- **BillingModule**: Handle pay-per-inference fee settlement
- **RewardModule**: Distribute rewards to node operators
- **SlashingModule**: Penalize misbehaving nodes
- **TokenModule**: MCELL token for staking and payments

## Getting Started

### Prerequisites

- Go 1.21+
- Cosmos SDK v0.50+

### Installation

```bash
# Clone the repository
git clone https://github.com/yueijdguedh/ReflexNet.git
cd ReflexNet

# Install dependencies
go mod download

# Build the binary
make build
```

## Project Status

🚧 Under active development

## License

MIT License

