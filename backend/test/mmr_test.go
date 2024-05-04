package main

import (
	"os"
	"testing"
)

func TestMMRCalculation(t *testing.T) {

}

func TestMain(m *testing.M) {

	// Run tests
	exitCode := m.Run()

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
