#!/bin/bash
################################################################################
Help()
{
   # Display Help
    echo "testService.sh is used to start (or stop) a local kubernetes instance for testing korwarder."
    echo
    echo "Syntax: testService.sh [-s|h|u|d]"
    
    echo "options:"
    echo "s     Attempt to start a test cluster locally"
    echo "h     Print this Help (meta!)."
    echo "u     Get status of korwarder local test cluster"
    echo "d     Delete local test cluster"
    echo
}
################################################################################
# Check for test deps - kind and kubectl
CheckDeps(){
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
    kubectl apply -f service1.yaml
    kubectl apply -f service2.yaml
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
    kubectl get services
    wait
      shift
      shift
      ;;
    # Goofy note about kind delete : It will print "..." if the cluster at spec. --name doesn't exist
    -d|--delete)
    CheckDeps
    kubectl delete -f service1.yaml
    kubectl delete -f service2.yaml
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
