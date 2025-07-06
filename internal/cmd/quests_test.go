package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestQuestsBrowseCommand(t *testing.T) {
	// Setup: create a sample quest directory and quest.yaml
	os.MkdirAll("quests/sample-quest", 0755)
	os.WriteFile("quests/sample-quest/quest.yaml", []byte("id: sample-quest\ntitle: Sample Quest\n"), 0644)
	defer os.RemoveAll("quests/sample-quest")

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"quests", "browse"})
	_ = rootCmd.Execute()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	if output != "sample-quest\n" {
		t.Errorf("expected 'sample-quest', got '%s'", output)
	}
}
