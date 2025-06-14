#!/bin/bash
VERSION=$1

if [ -z "$VERSION" ]; then
  echo "Usage: bash ./create-tag.sh v1.0.0"
  exit 1
fi

git pull
git tag "$VERSION" && git push origin "$VERSION"
