package command

import (
	"testing"
)

func Test_Execute(t *testing.T) {
	// commands for testing
	var inputRequired = map[string][]string{
		"input": []string{"-i", "the input file"},
	}
	var commands = map[string][]string{
		"input":     []string{},
		"output":    []string{"-o", "output"},
		"reverse":   []string{"-r"},
		"pick":      []string{"-p", "1", "-p", "4"},
		"name":      []string{"name"},
		"time":      []string{"time"},
		"size":      []string{"size"},
		"extension": []string{"extension"},
	}

	for k, v := range commands {
		cmd := NewCommand()
		args := append(inputRequired["input"], v...)
		cmd.SetArgs(args)
		config := cmd.Execute()
		t.Log(k, "\t->\t", *config)
	}
}
