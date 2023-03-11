#!/bin/bash
set -e
export SCRIPT_FILE=$(readlink -f "$0")
export SCRIPT_DIR=$(dirname "${SCRIPT_FILE}")
cd "${SCRIPT_DIR}"
buf generate
