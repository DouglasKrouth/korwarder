# korwarder 
# Dependencies
## User
Kubectl - latest [https://kubernetes.io/docs/tasks/tools/](kubectl download)
- TODO : pin a version
## Development only
Kind - latest [https://kind.sigs.k8s.io/](kind download)
- Used for running k8's 

### Notes on testing
Tests will need a way of simulating K8's for use in trying out port forward commands
- Idea is to use *kind* to implement a basic k8's image that has some exposed ports which we can run port forwards against for testing
