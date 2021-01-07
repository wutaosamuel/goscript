package job

import (
	"fmt"
	"os"

	"goscript/listnames"
	"goscript/listnames/config"
	c "goscript/renames/config"
)

// List display files orders
func (j *Job) List() {
	result := j.list(j.ListOperation)

	// print out result
	fmt.Println("List file name: ")
	fmt.Println()
	if len(j.Pick) == 0 {
		listnames.ConsoleString(result)
	}
	if len(j.Pick) != 0 {
		listnames.SelectConsoleString(result, j.SelectedFiles)
	}
}

// list sorted files
// 	- auto default
// 	- it will check picked files
//	- it will check if reverse order
func (j *Job) list(listOperation c.ListOperation) []string {
	result, err := listnames.Execute(j.Files, j.Pick, j.Reverse, convertListOperation(listOperation))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	selectedFiles := SelectFiles(result, j.Pick)

	// reverse file list order
	if j.Reverse {
		for i, j := 0, len(j.Files)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
		for i, j := 0, len(selectedFiles)-1; i < j; i, j = i+1, j-1 {
			selectedFiles[i], selectedFiles[j] = selectedFiles[j], selectedFiles[i]
		}
	}
	// set file list and selected file list
	j.Files = result
	j.SelectedFiles = selectedFiles

	return result
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
