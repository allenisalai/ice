package ice

import (
	"os"
	"testing"
)

func TestGetConfigFileLocation(t *testing.T) {
	oldHome := os.Getenv("HOME")
	defer func() {
		os.Setenv("HOME", oldHome)
	}()

	expectedHomeDir := "/home/test"
	os.Setenv("HOME", expectedHomeDir)

	expectedDir := expectedHomeDir + "/.ice/config.json"

	loc := getConfigFileLocation()

	if loc != expectedDir {
		t.Error("Configuration file didn't match expected filename")
	}
}
