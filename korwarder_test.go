package main

import "testing"

func TestKorwarderAddPortForward(t *testing.T) {
    k := Korwarder{}
    k.addPortForward(PortForward{"run a port forward command", "1", "1"})
    if len(k.pfs) != 1 {
        t.Errorf("len(addPf) = %d; want 1", len(k.pfs))
    }
}

func TestKorwarderAddPortForwardByCommand(t *testing.T) {
    k := Korwarder{}
    k.addPortForwardByCommand("run a port forward command", "1")
    k.addPortForwardByCommand("run a another port forward command", "2")
    if len(k.pfs) != 2 {
        t.Errorf("len(addPf) = %d; want 2", len(k.pfs))
    }
}
