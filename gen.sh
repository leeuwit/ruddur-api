#!/bin/bash
set -e

# Find the script's dir and cd to it
export SCRIPT_FILE=$(readlink -f "$0")
export SCRIPT_DIR=$(dirname "${SCRIPT_FILE}")
cd "${SCRIPT_DIR}"

# Generate code
buf mod update
buf lint
buf generate

# Post-process code
rm ruddur/v1/openapi.pb.go
sed -i 's|_ "google.golang.org/genproto/googleapis/api/annotations"||g' ruddur/v1/*.go
go mod tidy
