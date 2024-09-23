package main

import "testing"

func TestKorwarderAddPortForward(t *testing.T) {
	k := Korwarder{make(map[string]PortForward), "test"}
	k.addPortForward(PortForward{"run a port forward command"})
	if len(k.pfs) != 1 {
		t.Errorf("len(addPf) = %d; want 1", len(k.pfs))
	}
}

func TestKorwarderAddPortForwardByCommand(t *testing.T) {
	k := Korwarder{make(map[string]PortForward), "test"}
	k.addPortForwardByCommand("run a port forward command")
	k.addPortForwardByCommand("run a another port forward command")
	if len(k.pfs) != 2 {
		t.Errorf("len(addPf) = %d; want 2", len(k.pfs))
	}
}

func TestKorwarderGetPortForwards(t *testing.T) {
	k := Korwarder{make(map[string]PortForward), "test"}
	k.addPortForwardByCommand("test0")
	k.addPortForwardByCommand("test1")
	temp := k.getPortForwards()
	if len(temp) != 2 {
		t.Errorf("len(addPf) = %d; want 2", len(temp))
	}
	if temp[0].command != "test0" {
		t.Errorf("temp[0] = %v; want 'test0'", temp[0])
	}
}
