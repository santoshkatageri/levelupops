package cmd

import (
	"bytes"
	"testing"
)

// The TestCleanupCommand function tests the cleanup command output in a Go program.
func TestCleanupCommand(t *testing.T) {
	old := rootCmd.OutOrStdout()
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	rootCmd.SetArgs([]string{"cleanup"})
	_ = rootCmd.Execute()

	output := buf.String()
	expected := "Cleanup complete. (not implemented yet)\n"
	if output != expected {
		t.Errorf("expected '%s', got '%s'", expected, output)
	}

	rootCmd.SetOut(old)
}
