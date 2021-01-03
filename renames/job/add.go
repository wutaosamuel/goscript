package job

import (
	"fmt"
	"strconv"
	"os"
)

// Add add func
func (j *Job) Add() {
	// sort and list file orders
	j.list()

	// setup all files or selected files
	fileJob := NewFileJob()
	if len(j.SelectedFiles) == 0 {
		fileJob.ParseFilenames(j.Files)
	}
	if len(j.SelectedFiles) != 0 {
		fileJob.ParseFilenames(j.SelectedFiles)
	}

	char := j.Char
	number := j.Number
	if j.Number < 0 {
		number = 0
	}
	// add at end of file name
	if !j.AtBegin {
		for k, v := range fileJob.Files {
			v.Name += char
			if char != "" && j.Number == -1 {
				continue
			}
			v.Name += strconv.Itoa(number + k)
		}
	}
	// add at begin of file name
	if j.AtBegin {
		for k, v := range fileJob.Files {
			s := char
			if char != "" && j.Number == -1 {
				continue
			}
			s += strconv.Itoa(number + k)
			v.Name = s + v.Name
		}
	}

	// result files, it will check the same files
	result, contain := fileJob.ToFileNames(j.OutputDir)
	if contain {
		fmt.Println("Contains the same files, pls enter different char and number")
		os.Exit(0)
	}
	// add -> rename
	if j.OutputDir == "" {
		for k, v := range result {
			if k != fileJob.Files[k].ID {
				fmt.Println("Add rename Error")
				os.Exit(0)
			}
			os.Rename(fileJob.Files[k].ToFullFileName(), v)
		}
	}
	// add -> copy to target dir
	if j.OutputDir != "" {
		for k, v := range result {
			if k != fileJob.Files[k].ID {
				fmt.Println("Add copy error")
				os.Exit(0)
			}
			Copy(fileJob.Files[k].ToFullFileName(), v)
		}
	}
}
