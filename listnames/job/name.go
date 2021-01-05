package job

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// NameJob namejob object
type NameJob struct {
	ID int // file id
	Name string // full file name
	Base string // file base
}

// SortName sort files by name
// 	- ignore file extension
// 	- sort name only
func (j *Job) SortName() []string {
	var nameJobs []NameJob
	var result []string

	var filesLen = len(j.Files)
	nameJobs = make([]NameJob, 0, filesLen)
	result = make([]string, 0, filesLen)

	for k, v := range j.Files {
		extension := filepath.Ext(v)
		base := v[0 : len(v)-len(extension)]
		nameJobs = append(nameJobs, NameJob{k, v, base})
	}

	// sort by name: a, b, c, d ...
	if !j.Reverse {
		sort.Slice(nameJobs, func(i, j int) bool {
			return nameJobs[i].Base < nameJobs[j].Base
		})
	}
	// reverse sort by name: z, y, x ...
	if j.Reverse {
		sort.Slice(nameJobs, func(i, j int) bool {
			return nameJobs[i].Base > nameJobs[j].Base
		})
	}

	// result
	for _, v := range nameJobs {
		result = append(result, j.Files[v.ID])
	}

	return result
}

// checkExtension not allow different extension file
// TODO: support different files
func (j *Job) checkExtension(fileNames []string) {
	var extension string = ""

	for _, f := range fileNames {
		if extension == "" {
			extension = filepath.Ext(f)
			continue
		}
		if extension != filepath.Ext(f) {
			fmt.Println("Name operation not support different extension file now")
			os.Exit(0)
		}
	}
}
