#!/bin/bash
# Example: Register a new AI model

reflexnetd tx modelregistry register-model \
  --name="GPT-Style-Model" \
  --metadata-cid="QmYourMetadataCID123" \
  --shard-count=20 \
  --version="1.0.0" \
  --from=mykey \
  --chain-id=reflexnet-1 \
  --gas=auto \
  --yes

