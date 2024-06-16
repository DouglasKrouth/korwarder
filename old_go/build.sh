#!/bin/bash

# Lazy build script to get us started
mkdir -p ./bin
cd ./cmd
go build -o ../bin
