package main

import ("log"
        "os/exec"
)


func main() {
    checkKubectlInstalled()
}

func checkKubectlInstalled(){
    _, err := exec.Command("kubectl").Output()
    if err != nil {
        log.Fatal("kubectl is not installed on this machine and/or is inacessible at command 'kubectl'.\n1. Install it from https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/.\n2. Try again.")
    }
}
