#!/bin/bash
set -e
echo "Running tests..."
go test ./x/... -v
echo "Tests completed!"
