# ReflexNet API Documentation

## Overview

This document describes the REST and gRPC APIs for the ReflexNet protocol.

## Model Registry API

### Register Model

Register a new AI model in the system.

**Endpoint:** `POST /reflexnet/modelregistry/v1/register`

**Request:**
```json
{
  "owner": "reflex1...",
  "name": "MyModel",
  "metadata_cid": "QmXxx...",
  "shard_count": 10,
  "version": "1.0.0"
}
```

**Response:**
```json
{
  "model_id": 1
}
```

### Query Model

Query information about a specific model.

**Endpoint:** `GET /reflexnet/modelregistry/v1/models/{model_id}`

**Response:**
```json
{
  "model": {
    "id": 1,
    "owner": "reflex1...",
    "name": "MyModel",
    "metadata_cid": "QmXxx...",
    "shard_count": 10,
    "version": "1.0.0",
    "status": "MODEL_STATUS_ACTIVE",
    "inference_count": 0
  }
}
```

### List Models

List all registered models with pagination.

**Endpoint:** `GET /reflexnet/modelregistry/v1/models`

**Query Parameters:**
- `pagination.limit`: Number of results per page
- `pagination.offset`: Number of results to skip

**Response:**
```json
{
  "models": [...],
  "pagination": {
    "next_key": null,
    "total": 10
  }
}
```

## Shard Allocator API

### Register Node

Register a new shard node.

**Endpoint:** `POST /reflexnet/shardallocator/v1/register`

**Request:**
```json
{
  "node_address": "reflex1...",
  "stake_amount": {
    "denom": "mcell",
    "amount": "10000000"
  }
}
```

### Query Node Info

Query information about a specific node.

**Endpoint:** `GET /reflexnet/shardallocator/v1/nodes/{node_address}`

**Response:**
```json
{
  "node_info": {
    "address": "reflex1...",
    "staked_amount": "10000000",
    "total_shards": 5,
    "uptime_percentage": "0.98",
    "reputation_score": 100,
    "status": "NODE_STATUS_ACTIVE"
  }
}
```

## Inference Gateway API

### Submit Inference Request

Submit a new inference request.

**Endpoint:** `POST /reflexnet/inferencegateway/v1/inference`

**Request:**
```json
{
  "requester": "reflex1...",
  "model_id": 1,
  "input_data": "QmInputDataCID...",
  "nonce": 123
}
```

**Response:**
```json
{
  "request_id": "req_abc123..."
}
```

### Query Inference Status

Query the status of an inference request.

**Endpoint:** `GET /reflexnet/inferencegateway/v1/requests/{request_id}`

**Response:**
```json
{
  "request": {
    "request_id": "req_abc123...",
    "model_id": 1,
    "requester": "reflex1...",
    "status": "INFERENCE_STATUS_COMPLETED",
    "created_at": 12345
  }
}
```

## Billing API

### Query Billing Record

Query billing information for an inference request.

**Endpoint:** `GET /reflexnet/billing/v1/records/{request_id}`

**Response:**
```json
{
  "record": {
    "request_id": "req_abc123...",
    "model_id": 1,
    "requester": "reflex1...",
    "total_cost": "5000",
    "base_fee": "1000",
    "compute_fee": "4000",
    "compute_units": 100,
    "payment_status": "PAYMENT_STATUS_COMPLETED"
  }
}
```

## Reward API

### Claim Rewards

Claim accumulated rewards for a node.

**Endpoint:** `POST /reflexnet/reward/v1/claim`

**Request:**
```json
{
  "node_address": "reflex1..."
}
```

**Response:**
```json
{
  "claimed_amount": "15000"
}
```

### Query Node Rewards

Query reward information for a specific node.

**Endpoint:** `GET /reflexnet/reward/v1/rewards/{node_address}`

**Response:**
```json
{
  "reward": {
    "node_address": "reflex1...",
    "accumulated_rewards": "20000",
    "claimed_rewards": "5000",
    "performance_metrics": {
      "successful_inferences": 150,
      "failed_inferences": 2,
      "total_uptime_blocks": 10000
    }
  }
}
```

## Error Responses

All endpoints may return error responses in the following format:

```json
{
  "error": {
    "code": 3,
    "message": "model not found",
    "details": []
  }
}
```

## Common Error Codes

- `2`: Invalid argument
- `3`: Not found  
- `5`: Already exists
- `7`: Permission denied
- `13`: Internal error

