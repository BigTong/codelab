#!/bin/bash
protoc --go_out=plugins=grpc:./.. user.proto