// Methods to handle running the port-forward go-routines
package main

import (
    "fmt"
)

// represents a single instance of a `kubectl` port-forward command
type Pf struct {
    alive bool
    restarts int32
}

func (p *Pf) getStatus () {
    if p.alive == true {
        fmt.Println("It's alive!")
    } else {
        fmt.Println("It's dead!")
    }
    
}

// TODO : Will have multiple args passed
func PortForward(pfCommand string) {

    var temp Pf = Pf{alive: true, restarts: 3}
    temp.getStatus()
}
