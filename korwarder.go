package main

import "github.com/google/uuid"

// Represents a single port-forward object
type PortForward struct {
    command string
    group string
    ID string
}

// Contains available port-forwards, relationships at runtime
type Korwarder struct {
    // TODO : Need to convert this into something like map[pf_uuid][]PortForward
    // TODO : update PortForward object to have uuid passed as param, pass uuid as param
    pfs []PortForward
}

func (k *Korwarder) addPortForward(toAdd PortForward) []PortForward {
    k.pfs = append(k.pfs, toAdd)
    return k.pfs
}

func (k *Korwarder) addPortForwardByCommand(command string, group string) []PortForward {
    k.pfs = append(k.pfs, PortForward{command, group, uuid.New().String()})
    return k.pfs
}

// TODO : Needs to be implemented
func (k *Korwarder) removePortForwardByID(toRemove PortForward) []PortForward {
    return k.pfs
}
