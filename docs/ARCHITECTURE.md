# ReflexNet Architecture

## System Overview

ReflexNet (MindCell) is a decentralized protocol built on Cosmos SDK for AI model sharding, inference execution, and cryptoeconomic incentives. The system enables secure, verifiable, and monetizable AI inference through blockchain technology.

## Core Components

### 1. Model Registry Module

**Purpose:** Manages registration, versioning, and metadata of AI models.

**Key Features:**
- Model registration with IPFS/Arweave metadata storage
- Version control for model updates
- Owner permission management
- Status tracking (Active, Inactive, Deprecated)

**State:**
- Models indexed by ID
- Models indexed by owner
- Module parameters (registration fee, max metadata size, max shard count)

### 2. Shard Allocator Module

**Purpose:** Manages the assignment and distribution of model shards across validator nodes.

**Key Features:**
- Node registration and staking
- Shard assignment algorithms
- Load balancing across nodes
- Reputation-based allocation
- Health check monitoring

**State:**
- Shard assignments (model ID + shard index → node address)
- Node information (stake, reputation, uptime)
- Node status tracking

**Algorithms:**
- **Load Balancing:** Distributes shards evenly across available nodes
- **Reputation Weighting:** Nodes with higher reputation get priority
- **Health Monitoring:** Regular checks ensure node availability

### 3. Inference Gateway Module

**Purpose:** Routes inference requests and manages zkML proof verification.

**Key Features:**
- Request routing to shard nodes
- Response aggregation
- zkML proof generation and verification
- Nonce-based replay protection
- Rate limiting per address

**Workflow:**
1. Client submits inference request with input data
2. Gateway routes request to nodes holding required shards
3. Nodes compute partial results off-chain
4. zkML prover aggregates results and generates proof
5. Gateway verifies proof on-chain
6. Upon verification, triggers billing and reward distribution

**State:**
- Inference requests and their status
- Inference responses with proofs
- Request nonces for replay protection

### 4. Billing Module

**Purpose:** Handles pay-per-inference fee calculation and distribution.

**Key Features:**
- Dynamic pricing based on compute units
- Multi-party fee distribution
- Payment status tracking

**Fee Structure:**
```
total_cost = base_fee + (compute_unit_price * compute_units)
```

**Distribution:**
- Model Owner: Configurable percentage (e.g., 40%)
- Shard Nodes: Configurable percentage (e.g., 50%)
- Protocol Treasury: Configurable percentage (e.g., 10%)

**State:**
- Billing records per inference request
- Payment distributions
- Module parameters (base fee, compute unit price, split percentages)

### 5. Reward Module

**Purpose:** Manages incentives and penalties for node operators.

**Key Features:**
- Periodic reward distribution
- Performance-based multipliers
- Slashing for misbehavior
- Reward pool management

**Reward Calculation:**
```
node_reward = base_reward * performance_multiplier * uptime_factor
```

**Slashing Events:**
- Incorrect zkML proof submission
- Timeout in responding to inference requests
- Node unavailability during health checks
- Detected malicious behavior

**State:**
- Reward pool with total and distributed amounts
- Node rewards (accumulated and claimed)
- Performance metrics per node
- Slashing event history

## Data Flow

### Inference Request Flow

```
1. Client → InferenceGateway: Submit request with input data
2. InferenceGateway → ShardAllocator: Query shard assignments
3. InferenceGateway → Shard Nodes: Distribute compute tasks
4. Shard Nodes → zkML Prover: Submit partial results
5. zkML Prover → InferenceGateway: Return proof and aggregated result
6. InferenceGateway: Verify proof on-chain
7. InferenceGateway → Billing: Trigger payment
8. Billing → Reward: Distribute rewards to nodes
9. InferenceGateway → Client: Return verified result
```

### Node Lifecycle

```
1. Registration: Node stakes tokens and registers
2. Shard Assignment: ShardAllocator assigns model shards
3. Active Service: Node responds to inference requests
4. Health Checks: Periodic verification of availability
5. Reward Distribution: Periodic rewards based on performance
6. Unregistration: Node unstakes and exits
```

## Security Mechanisms

### 1. zkML Verification

Zero-knowledge proofs ensure that inference results are mathematically correct without revealing the full model. This prevents:
- Forged inference outputs
- Model theft through reverse engineering
- Result manipulation

### 2. Staking and Slashing

Economic security through:
- Required stake for node participation
- Slashing for provably incorrect behavior
- Reputation system for long-term reliability

### 3. Replay Protection

- Nonce-based request identification
- Prevents duplicate inference charges
- Ensures request uniqueness

### 4. Rate Limiting

- Per-address request limits
- Prevents DoS attacks
- Ensures fair resource access

### 5. Metadata Integrity

- IPFS/Arweave content addressing
- Ensures immutable model metadata
- Verifiable through CID references

## Scalability Considerations

### Horizontal Scaling

- Multiple nodes can serve different shards
- Load distributed across the network
- New nodes can join dynamically

### State Management

- Efficient key-value store design
- Indexed queries for fast lookups
- Pagination for large result sets

### Off-Chain Computation

- Inference computed off-chain
- Only proofs verified on-chain
- Reduces blockchain bloat

## Integration Points

### IPFS/Arweave

- Store model metadata
- Store large input/output data
- Reference via CID in on-chain transactions

### zkML Systems

- Proof generation (gnark, Halo2, custom circuits)
- Verification on-chain
- Public inputs management

### IBC (Future)

- Cross-chain inference requests
- Multi-chain model registry
- Interoperable reward systems

## Module Dependencies

```
ModelRegistry ← ShardAllocator ← InferenceGateway
                      ↓                 ↓
                    Reward ← Billing ←─┘
```

- ShardAllocator depends on ModelRegistry for model information
- InferenceGateway depends on both ModelRegistry and ShardAllocator
- Billing depends on InferenceGateway for completed requests
- Reward depends on Billing for payment information and ShardAllocator for node info

## Future Enhancements

1. **FHE Integration:** Privacy-preserving inference using fully homomorphic encryption
2. **Cross-Chain Support:** IBC-enabled inference requests from other chains
3. **Model NFTs:** Tokenized model ownership and licensing
4. **Dataset Watermarking:** On-chain verification of dataset attribution
5. **Adaptive Pricing:** Dynamic fee adjustment based on network demand

