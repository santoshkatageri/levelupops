package cmd

import (
	"bytes"
	"testing"
)

// The TestDoctorCommand function tests the doctor command output in a Go program.
func TestDoctorCommand(t *testing.T) {
	old := rootCmd.OutOrStdout()
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	rootCmd.SetArgs([]string{"doctor"})
	_ = rootCmd.Execute()

	output := buf.String()
	expected := "Diagnostics passed. (not implemented yet)\n"
	if output != expected {
		t.Errorf("expected '%s', got '%s'", expected, output)
	}

	rootCmd.SetOut(old)
}
