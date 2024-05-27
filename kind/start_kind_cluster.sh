#!/bin/bash
# Script to start a kind cluster that will be used for local integration testing `korwarder`.
# Author : Douglas Krouth

if ! command -v asdfadsfkind &> /dev/null
then
    printf "kind could not be found.\nInstall it at https://kind.sigs.k8s.io/docs/user/quick-start/ and rerun this script.\n"
    exit 1
fi
# kind create cluster --config config.yaml
