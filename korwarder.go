package main

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

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

// Whether or not we have any
func (k *Korwarder) checkIfStorageExists() bool {
	var storagePath = filepath.Join(k.appDataPath, "data.json")
	_, err := os.Stat(storagePath)
	if os.IsNotExist(err) {
		log.Debug().Msgf("No stored korwarder data found at %s", storagePath)
		return false
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO : this needs to be fleshed out
func (k *Korwarder) loadStorage() bool {
	var storagePath = filepath.Join(k.appDataPath, "data.json")
	if k.checkIfStorageExists() {
		data, err := os.ReadFile(storagePath)
		check(err)
		log.Printf("%v\n", data)
	} else {
		// TODO : implement?
		log.Debug().Msgf("No .korwarder storage found at %s", storagePath)
		return false
	}
	return true
}

func (k *Korwarder) addPortForward(toAdd PortForward) map[string]PortForward {
	k.pfs[uuid.New().String()] = toAdd
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
