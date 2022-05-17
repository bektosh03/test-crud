#!/bin/bash
CURRENT_DIR=$1
PROTO_PATH=$CURRENT_DIR
PREFIX="github.com/bektosh03/test-crud/genprotos"

./.local/bin/protoc --proto_path="${PROTO_PATH}"/ --go_opt=module="${PREFIX}" \
    --go_out="${CURRENT_DIR}" --go-grpc_out="${CURRENT_DIR}" --go-grpc_opt=module="${PREFIX}"\
    "${CURRENT_DIR}"/protos/datapb/*.proto