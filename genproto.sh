#!/bin/bash

set -e
set -o pipefail
set -x

if [[ -z "${GITHUB_WORKSPACE}" ]]; then
  GIT_WORK_DIR=$(git rev-parse --show-toplevel)
else
  GIT_WORK_DIR="${GITHUB_WORKSPACE}"
fi

(cd $GIT_WORK_DIR && protoc --go_opt=module=go.verytable.online/aot --go_out=. $(find ./aot/proto/client/api/rpc -iname "*.proto"))

