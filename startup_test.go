package main

import (
	"os"
	"testing"
)

func TestSetupDataFolderHomeDir(t *testing.T) {
	t.Setenv("KORWARDER_APP_DATA_PATH", "testdata")
	setupAppDataFolder()
	os.RemoveAll("./testdata/")
}
