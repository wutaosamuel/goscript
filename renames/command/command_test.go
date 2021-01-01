package command

import (
	"testing"
)

func Test_Execute(t *testing.T) {
	// commands for testing
	var inputRequired = map[string][]string{
		"input": []string{"-i", "./"},
	}
	var commands = map[string][]string{
		"input":   []string{},
		"output":  []string{"-o", "./"},
		"begin":   []string{"-b"},
		"pick":    []string{"-p", "1", "-p", "4"},
		"reverse": []string{"-r"},
		"list":    []string{"-l", "1"},
		"add":     []string{"add", "-c", "added char"},
		"delete":  []string{"delete", "-n", "5"},
		"rename":  []string{"rename", "-c", "renamed char"},
		"listcom": []string{"list", "-e"},
		"count":   []string{"count", "-c", "counted char"},
	}

	for k, v := range commands {
		cmd := NewCommand()
		args := append(inputRequired["input"], v...)
		cmd.SetArgs(args)
		config := cmd.Execute()
		t.Log(k, " -> ", config.Files, config.OutputDir,
			config.Begin, config.Pick, config.Reverse,
			config.ListFile, config.Add, config.Delete,
			config.Rename, config.List, config.Count)
	}
}
