# korwarder (maybe rename to "korridor")?

**Goal** : Create a simple-to-use, out-of-cluster, terminal based ui (and TUI?) that simplifies using port-forwards from the command line. This project does not do anything fancy with the Kubernetes API/SDK, just calls `kubectl` directly to handle any `port-forward` operations.

**Core functionality**
1. Ability to save, edit, delete port-forward commands
2. Check running port-forwards
3. Ability to add comments/data about port-forward commands (example. "#dev service", etc.)
4. Ability to group multiple port-forward commands together, run all of them with a single command.

**Extra functionality**
1. Have the ability to change (automatically?) the list of port-forwards/commands based on the cluster context.


**Components**
- forward module : handles start/stop for goroutines to run pf's, restart/retry policy
- tui module : Actual application, contains main
