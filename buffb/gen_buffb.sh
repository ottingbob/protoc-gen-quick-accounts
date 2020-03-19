#!/bin/sh

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  --go_out=plugins=grpc:. \
  --swagger_out=logtostderr=true,allow_delete_body=true:../docs/ \
  --grpc-gateway_out=logtostderr=true,allow_delete_body=true:. \
  ./*.proto
