package job

import (
	"fmt"
	"os"
)

// Delete delete func
func (j *Job) Delete() {
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

	// delete to file names
	number := j.Number
	if j.Number == 0 {
		return
	}
	if j.Number < 0 {
		number = 1
	}
	// delete at end of filename
	if !j.AtBegin {
		for _, v := range fileJob.Files {
			if len(v.Name) <= number {
				fmt.Println("Not allow to delete too much.")
				os.Exit(0)
			}
			v.Name = v.Name[0 : len(v.Name)-number]
		}
	}
	// delete at begin of filename
	if j.AtBegin {
		for _, v := range fileJob.Files {
			if len(v.Name) <= number {
				fmt.Println("Not allow to delete too much.")
				os.Exit(0)
			}
			v.Name = v.Name[number:0]
		}
	}

	// result files, it will check the same files
	result, contain := fileJob.ToFileNames(j.OutputDir)
	if contain {
		fmt.Println("Contains the same files, pls enter different char and number")
		os.Exit(0)
	}
	// delete -> rename
	if j.OutputDir == "" {
		for k, v := range result {
			if k != fileJob.Files[k].ID {
				fmt.Println("Add rename Error")
				os.Exit(0)
			}
			os.Rename(files[k], v)
		}
	}
	// delete -> copy to target dir
	if j.OutputDir != "" {
		for k, v := range result {
			if k != fileJob.Files[k].ID {
				fmt.Println("Add rename Error")
				os.Exit(0)
			}
			Copy(files[k], v)
		}
	}
}
