#!/bin/bash
# Description: Builds the protocolbuffers into a go package

echo -e "\033[4mCompiling protocol buffers...\033[0m"
command -v protoc >/dev/null 2>&1 || { echo "'protoc' tool is required but missing." >&2; exit 1; }
command -v protoc-gen-go >/dev/null 2>&1 || { echo "'protoc-gen-go' tool is required but missing." >&2; exit 1; }

#Clean up temp dir
rm -rf ./tmp_proto

#Clone down the latest proto
(git clone https://github.com/webhookd/api-proto.git tmp_proto && cd ./tmp_proto)

#Create the output directory
mkdir -p webhookd

#Generate Output
protoc -I ./tmp_proto --gogo_out=plugins=grpc,paths=source_relative:./webhookd ./tmp_proto/*.proto --proto_path=$GOPATH/src

#Remove the temp dir
rm -rf ./tmp_proto