#!/bin/bash
# Simple script to delete the local kind cluster (which uses config.yaml).
# Author : Douglas Krouth

export CLUSTER_NAME=$(grep "^name\:" config.yaml | sed 's/^.*: //')
kind delete clusters $CLUSTER_NAME
