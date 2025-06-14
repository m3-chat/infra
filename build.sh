#!/usr/bin/env bash

set -e

echo "Tidying and downloading modules..."
go mod tidy
go mod download

echo "Installing CLI..."
go install ./cmd/infra

echo "Building binary..."
go build -o infra ./cmd/infra

echo "✅ Build complete: ./infra"
