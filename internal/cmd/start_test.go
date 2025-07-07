package cmd

import (
	"bytes"
	"os"
	"testing"
)

// The TestStartCommand_Success function tests the start command functionality by creating a test
// quest, executing the command, and checking the output message.
func TestStartCommand_Success(t *testing.T) {
	os.MkdirAll("quests/test-quest", 0755)
	os.WriteFile("quests/test-quest/quest.yaml", []byte("id: test-quest\ntitle: Test Quest\n"), 0644)
	defer os.RemoveAll("quests/test-quest")

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"start", "test-quest"})
	_ = rootCmd.Execute()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	expected := "Quest 'test-quest' started!\n"
	if output != expected {
		t.Errorf("expected '%s', got '%s'", expected, output)
	}
}

// The TestStartCommand_NotFound function tests the output when a specific quest is not found.
func TestStartCommand_NotFound(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"start", "notfound"})
	_ = rootCmd.Execute()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	expected := "Quest 'notfound' not found.\n"
	if output != expected {
		t.Errorf("expected '%s', got '%s'", expected, output)
	}
}
