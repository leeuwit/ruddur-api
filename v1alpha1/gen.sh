#!/bin/bash
set -e
export SCRIPT_FILE=$(readlink -f "$0")
export SCRIPT_DIR=$(dirname "${SCRIPT_FILE}")
cd "${SCRIPT_DIR}"
protoc -I. \
  -Ivendor \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=. \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt generate_unbound_methods=true \
  --openapiv2_out=. \
  --openapiv2_opt logtostderr=true \
  ruddur.proto
