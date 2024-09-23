package main

import (
	"os"
	"testing"
)

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

func TestCheckIfStorageExists(t *testing.T) {
	k0 := Korwarder{make(map[string]PortForward), "test"}
	k1 := Korwarder{make(map[string]PortForward), "./testdata/.korwarder"}
	os.MkdirAll("./testdata/.korwarder", os.ModePerm)
    os.OpenFile("./testdata/.korwarder/data.json", os.O_RDONLY|os.O_CREATE, 0666)
	if k0.checkIfStorageExists() == true {
		t.Errorf("There shouldn't be any stored data at 'test`")
	}
	if k1.checkIfStorageExists() == false {
		t.Errorf("There should be a .korwarder/data.json file at `./testdata/.korwarder/data.json`")
	}
    os.RemoveAll("./testdata")
}

func TestLoadStorage(t *testing.T) {
	k0 := Korwarder{make(map[string]PortForward), "./test/.korwarder"}
	k1 := Korwarder{make(map[string]PortForward), "./testdata/.korwarder"}
	os.MkdirAll("./testdata/.korwarder", os.ModePerm)
    os.OpenFile("./testdata/.korwarder/data.json", os.O_RDONLY|os.O_CREATE, 0666)
    if k0.loadStorage() == true {
		t.Errorf("There should not be stored data at 'test/.korwarder`.")
    }
    if k1.loadStorage() == false {
		t.Errorf("There should be stored data at 'testdata/.korwarder` (it's just an empty 'data.json' for this test).")
    }
    os.RemoveAll("./testdata")
}
