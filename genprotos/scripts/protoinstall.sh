#!/bin/bash
CURRENT_DIR=$1
DEPENDENCIES_DIR_NAME=".local"
PROTOC_VERSION="3.20.1"
PROTOC_BIN_NAME="protoc"
PROTOC_BIN_PATH="$CURRENT_DIR/$DEPENDENCIES_DIR_NAME/bin/"
PB_REL="https://github.com/protocolbuffers/protobuf/releases"

if test -f "$PROTOC_BIN_PATH/$PROTOC_BIN_NAME"; then
    echo "found $PROTOC_BIN_NAME in $PROTOC_BIN_PATH. Skipping installation"
else
    echo "did not found $PROTOC_BIN_NAME in $PROTOC_BIN_PATH. Proceeding to installation..."
    curl -LO "$PB_REL/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip"
    unzip "protoc-$PROTOC_VERSION-linux-x86_64.zip" -d "$CURRENT_DIR/$DEPENDENCIES_DIR_NAME"
    rm "protoc-$PROTOC_VERSION-linux-x86_64.zip"
fi

echo "Installing protoc go dependencies..."

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

echo "Successfully finished!"