package job

import (
	"fmt"
	"os"
	"strconv"
)

// Add add func
func (j *Job) Add() {
	// sort and list file orders
	j.list(j.DoList)

	// setup all files or selected files
	fileJob := NewFileJob()
	files := make([]string, 0)
	if len(j.SelectedFiles) == 0 {
		files = j.Files
	}
	if len(j.SelectedFiles) != 0 {
		files = j.SelectedFiles
	}
	fileJob.ParseFilenames(files)

	// add to file name
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
	// add -> rename or move files
	for k, v := range result {
		if k != fileJob.Files[k].ID {
			fmt.Println("Add rename Error")
			os.Exit(0)
		}
		os.Rename(files[k], v)
	}
}
