#!/usr/bin/env bash

protoc -I ./ --go_out=plugins=grpc:../../source/rpc example.proto

protoc -I ./ --grpc-gateway_out=logtostderr=true:../../source/rpc example.proto
