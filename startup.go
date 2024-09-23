package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"path/filepath"
)

func startup() string {
	checkKubectlInstalled()
	var appDataPath = setupAppDataFolder()
	log.Printf("%s", appDataPath)
	return appDataPath
}

func checkKubectlInstalled() {
	_, err := exec.Command("kubectl").Output()
	if err != nil {
		// log.Fatal("kubectl is not installed on this machine and/or is inacessible at command 'kubectl'.\n1. Install it from https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/.\n2. Try again.")
		log.Fatal().Msg("kubectl is not installed on this machine and/or is inacessible at command 'kubectl'.\n1. Install it from https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/.\n2. Try again.")
	}
}

// TODO : This is a bit clunky, could probably be cleaned up
func setupAppDataFolder() string {
	// Check if we've defined the app data path in env; if not, use default /home/<USER>/.korwarder
	var appDataPath = ""
	var envAppDataPath = os.Getenv("KORWARDER_APP_DATA_PATH")
	if envAppDataPath != "" {
		appDataPath = filepath.Join(envAppDataPath, ".korwarder")
		_, err := os.Stat(appDataPath)
		// Case when app data path is defined in env, not created yet
		if err != nil {
			err = os.MkdirAll(appDataPath, os.ModePerm)
			if err != nil {
				log.Fatal().Msgf("Attempted to create file at %s, was unsuccessful.", appDataPath)
			}
			log.Info().Msgf("Successfully created korwarder app data at `%s`", appDataPath)
		}
	} else {
		appDataPath, err := os.UserHomeDir()
		if err != nil {
			log.Fatal().Msg("Unable to access user home dir")
		}
		appDataPath = filepath.Join(appDataPath, ".korwarder")
		os.MkdirAll(appDataPath, os.ModePerm)
		if err != nil {
			log.Fatal().Msgf("Attempted to create file at %s, was unsuccessful.", appDataPath)
		}
		log.Info().Msgf("Successfully created korwarder app data at `%s`", appDataPath)
	}

	return appDataPath
}
