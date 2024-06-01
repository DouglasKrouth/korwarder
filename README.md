# korwarder (maybe rename to "korridor")?

**Goal** : Create a simple-to-use, out-of-cluster, terminal based ui (TUI) that simplifies using port-forwards from the command line.
**Core functionality**
- User selects which port-forward(s) they want to run. 
- Each port-forward runs as a single `goroutine`; need functionality to handle port-forward errors on each `goroutine` (retry `n` times, stop if failure, etc.). Port-forwards should operate as distinct channels.

**Components**
- forward module : handles start/stop for goroutines to run pf's, restart/retry policy
- korwarder module : Actual application, contains main

**Stack**
Language : Golang, bash
Main deps
* client-go : https://github.com/kubernetes/client-go
  * For handling k8's work
* bubbletea : https://github.com/charmbracelet/bubbletea
  * For TUI/UI in terminal. Need to investigate zsh compatibility, probably won't support non-POSIX compliant terms immediately.
* bubblezone : (maybe) use this for clickable stuff

**Potential Enhancements**
1. Have configuration in place to recognize which kube context you're currently in, adjust accordingly.
2. Add goroutine logic to enable multiple threads (port-forwards) to run simultaneously.
3. General UI QOL stuff (TBD)
