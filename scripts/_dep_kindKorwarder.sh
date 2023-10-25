#!/bin/bash
################################################################################
Help()
{
   # Display Help
    echo "kindKorwarder is used to start (or stop) a kubernetes in docker (kind) instance for testing korwarder."
    echo
    echo "Syntax: kindKorwarder [-s|h|u|d]"
    
    echo "options:"
    echo "s     Attempt to start a kind cluster with name "kind-korwarder""
    echo "h     Print this Help (meta!)."
    echo "u     Get status of kind-korwarder local cluster"
    echo "d     Delete the kind-korwarder local cluster"
    echo
}
################################################################################
# Check for test deps - kind and kubectl
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
################################################################################
# Case for no args, return help
[ $# -eq 0 ] && { Help; exit; }

while [[ $# -gt 0 ]]; do
  case $1 in
    -s|--start)
    CheckDeps
    kind create cluster --name korwarder
    wait
    if [ $? -eq 1 ]; then
        echo "korwarder test k8's deployment already exists"
    fi
    # Shifts aren't really needed as we only take one arg but I like having them there
    shift # past argument
    shift # past value
      ;;
    -u|--status)
    CheckDeps
    kubectl cluster-info --context kind-korwarder
    wait
      shift
      shift
      ;;
    # Goofy note about kind delete : It will print "..." if the cluster at spec. --name doesn't exist
    -d|--delete)
    CheckDeps
    kind delete cluster --name korwarder
    wait
      shift
      shift
      ;;
    -h|--help)
    Help
      shift
      shift
      ;;
    -*|--*)
      echo "Unknown option $1"
      exit 1
      ;;
    *)
      POSITIONAL_ARGS+=("$1") # save positional arg
      shift # past argument
      ;;
  esac
done

exit 0
