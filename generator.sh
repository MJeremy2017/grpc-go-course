#!/usr/bin/env bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/protobuf/calculator.proto --go_out=plugins=grpc:.