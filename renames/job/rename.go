package job

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Rename delete func
func (j *Job) Rename() {
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

	// rename file names
	char := j.Char
	extension := j.Extension
	number := j.Number
	// do nothing
	if char == "" && number == -1 && extension == "" {
		return
	}
	if j.Number < 0 {
		number = 0
	}
	// delete at end of filename
	if !j.AtBegin {
		for k, v := range fileJob.Files {
			v.Name = char + strconv.Itoa(number+k)
			if extension != "" {
				v.Extension = "." + extension
			}
		}
	}
	// delete at begin of filename
	if j.AtBegin {
		for k, v := range fileJob.Files {
			v.Name = strconv.Itoa(number+k) + char
			if extension != "" {
				v.Extension = "." + extension
			}
		}
	}

	// result files, it will check the same files
	result, contain := fileJob.ToFileNames(j.OutputDir)
	if contain {
		fmt.Println("Contains the same files, pls enter different char and number")
		os.Exit(0)
	}
	// rename -> rename
	if j.OutputDir == "" {
		for k, v := range result {
			if k != fileJob.Files[k].ID {
				fmt.Println("Add rename Error")
				os.Exit(0)
			}
			os.Rename(files[k], v)
		}
	}
	// rename -> copy to target dir
	if j.OutputDir != "" {
		wg := new(sync.WaitGroup)
		for k, v := range result {
			if k != fileJob.Files[k].ID {
				fmt.Println("Add rename Error")
				os.Exit(0)
			}
			wg.Add(1)
			GoCopy(files[k], v, wg)
		}
		wg.Wait()
	}
}
