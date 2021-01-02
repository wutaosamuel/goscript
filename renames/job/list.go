package job

import (
	"fmt"
	"os"

	"../../listnames"
	"../../listnames/config"
	c "../config"
)

// List display files orders
func (j *Job) List() {
	result, err := listnames.Execute(j.Files, j.Pick, j.Reverse, convertListOperation(j.ListOperation))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print out result
	fmt.Println("List file name: ")
	fmt.Println()
	if len(j.Pick) == 0 {
		listnames.ConsoleString(result)
	}
	if len(j.Pick) != 0 {
		selectFiles := listnames.SelectFiles(result, j.Pick)
		listnames.SelectConsoleString(result, selectFiles)
	}
}

// convertListOperation convert Operation
func convertListOperation(operation c.ListOperation) config.OpCode {
	opInt := c.ListOperationInt(operation)

	if opInt < 0 || opInt > 4 {
		return config.ErrorOp
	}
	switch {
	case opInt == 0:
		return config.DefaultOp
	case opInt == 1:
		return config.NameOp
	case opInt == 2:
		return config.TimeOp
	case opInt == 3:
		return config.SizeOp
	case opInt == 4:
		return config.ExtensionOp
	}

	return config.ErrorOp
}
