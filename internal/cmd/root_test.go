package cmd

import (
	"bytes"
	"os"
	"testing"
)

// The TestInitCommand function tests the output of the rootCmd when initialized with the "init"
// command.
func TestInitCommand(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"init"})
	_ = rootCmd.Execute()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	if output != "LevelUpOps initialized!\n" {
		t.Errorf("expected 'LevelUpOps initialized!', got '%s'", output)
	}
}
