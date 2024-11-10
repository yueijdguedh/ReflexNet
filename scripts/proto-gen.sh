#!/bin/bash

# Generate protobuf files for ReflexNet

set -e

echo "Generating protobuf files..."

cd proto

# Run buf generate
buf generate

echo "Protobuf generation completed successfully!"

cd ..

# Format generated Go files
echo "Formatting Go files..."
go fmt ./...

echo "Done!"

