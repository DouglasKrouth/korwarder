package main

import "github.com/google/uuid"

// Represents a single port-forward object
type PortForward struct {
	command string
	// TODO : implement groups
	// groups []string
}

// Contains available port-forwards, relationships at runtime
type Korwarder struct {
	pfs         map[string]PortForward
	appDataPath string
}

func (k *Korwarder) addPortForward(toAdd PortForward) map[string]PortForward {
	k.pfs[uuid.New().String()] = toAdd
	// k.pfs = append(k.pfs, toAdd)
	return k.pfs
}

func (k *Korwarder) addPortForwardByCommand(command string) map[string]PortForward {
	k.pfs[uuid.New().String()] = PortForward{command}
	return k.pfs
}

func (k *Korwarder) getPortForwards() []PortForward {
	x := make([]PortForward, 0, len(k.pfs))

	for _, value := range k.pfs {
		x = append(x, value)
	}
	return x
}

// // TODO : Needs to be implemented
// func (k *Korwarder) removePortForwardByID(toRemove PortForward) []PortForward {
//     return k.pfs
// }
