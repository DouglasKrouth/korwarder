# korwarder (maybe rename to "korridor")?

**Goal** : Create a simple-to-use, out-of-cluster, terminal based ui (TUI) that simplifies using port-forwards from the command line.

**Stack**
Language : Golang, bash
Main deps
* client-go : https://github.com/kubernetes/client-go
  * For handling k8's work
* bubbletea : https://github.com/charmbracelet/bubbletea
  * For TUI/UI in terminal. Need to investigate zsh compatibility, probably won't support non-POSIX compliant terms immediately.

**Potential Enhancements**
1. Have configuration in place to recognize which kube context you're currently in, adjust accordingly.
2. Add goroutine logic to enable multiple threads (port-forwards) to run simultaneously.
3. General UI QOL stuff (TBD)
