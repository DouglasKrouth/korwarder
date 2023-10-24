#!/bin/bash
################################################################################
Help()
{
   # Display Help
    echo "kindKorwarder is used to start (or stop) a kubernetes in docker (kind) instance for testing korwarder."
    echo
    echo "Syntax: kindKorwarder [-s|h|u|d]"
    
    echo "options:"
    echo "s     Start a kind cluster with name "kind-korwarder""
    echo "h     Print this Help (meta!)."
    echo "u     Get status of kind-korwarder local cluster"
    echo "d     Delete the kind-korwarder local cluster"
    echo
}
################################################################################
# Check for test deps
CheckDeps(){
if ! command -v kind &> /dev/null
then
    echo "kind was not found on your machine. Please install the dependency https://kind.sigs.k8s.io/"
    exit 1
fi
if ! command -v kubectl &> /dev/null
then
    echo "kubectl was not found on your machine. Please install the dependency https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/"
    exit 1
fi
}

# Case for no args, return help
[ $# -eq 0 ] && { Help; exit; }

# Case for -h
while getopts ":h" option; do
   case $option in
      h) # display Help
         Help
         exit;;
   esac
   exit 0
done


kind create cluster --name korwarder
if [ $? -eq 1 ]; then
    echo "korwarder test k8's deployment already exists"
fi
Help
