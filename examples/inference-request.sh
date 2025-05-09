#!/bin/bash
# Submit an inference request
reflexnetd tx inferencegateway submit-request \
  --model-id=1 \
  --input-data="QmInputCID" \
  --from=mykey \
  --yes
