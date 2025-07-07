package cmd

import (
	"bytes"
	"testing"
)

// The TestProgressCommand function tests the output of the "progress" command in a Go application.
func TestProgressCommand(t *testing.T) {
	old := rootCmd.OutOrStdout()
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	rootCmd.SetArgs([]string{"progress"})
	_ = rootCmd.Execute()

	output := buf.String()
	expected := "XP: 0\nJournal: (not implemented yet)\n"
	if output != expected {
		t.Errorf("expected '%s', got '%s'", expected, output)
	}

	rootCmd.SetOut(old)
}
