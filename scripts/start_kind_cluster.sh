#!/bin/bash
# Script to start a kind cluster that will be used for local integration testing `korwarder`.
# Author : Douglas Krouth

# Check whether we've already deployed the mongodb pods. Naive, only checks if a single pod is up.
function check_deployed_pods () {
    if [[ $(kubectl get pods | grep mongo | awk '{print $1}') != "" ]]; then
        return 0
    else 
        return 1
    fi
}

# Inspiration from k8's docs. We just need a simple service that we can ping to test that the port-forward calls work
function deploy_pods () {
    kubectl apply -f https://k8s.io/examples/application/mongodb/mongo-deployment.yaml
    kubectl apply -f https://k8s.io/examples/application/mongodb/mongo-service.yaml
    kubectl get service mongo
}

if ! command -v kind &> /dev/null
then
    printf "kind command could not be found.\n1. Install it at https://kind.sigs.k8s.io/docs/user/quick-start/\n2. Re-run this script.\n"
    exit 1
else 
    resp="$(kind create cluster --config config.yaml)"
    if [[ $? -eq 1 ]]; then
        echo "Kind cluster already deployed."
    fi

    if [[ check_deployed_pods == 0 ]]; then
        echo "Mongo test pod deployed, exiting."
        exit 0
    else
        deploy_pods
        exit 0
    fi
    exit 1
fi

