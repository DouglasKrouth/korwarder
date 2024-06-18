# korwarder (maybe rename to "korridor")?

**Goal** : Create a simple-to-use, out-of-cluster, terminal based ui (TUI) that simplifies using port-forwards from the command line.
**Core functionality**
- User selects which port-forward(s) they want to run. 
- Tui interaction shows existing (live) port-forwards currently running on the machine.

**Components**
- forward module : handles start/stop for goroutines to run pf's, restart/retry policy
- korwarder module : Actual application, contains main

**Stack**
Language : Rust, bash
Main deps
    - Ratatui

**Potential Enhancements**
1. Have configuration in place to recognize which kube context you're currently in, adjust accordingly.
