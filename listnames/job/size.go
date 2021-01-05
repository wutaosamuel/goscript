package job

import (
	"os"
	"sort"
)

// SizeJob size job
type SizeJob struct {
	Name string // file name
	Size int64  // file size
}

// SortSize sort file name by time
func (j *Job) SortSize() []string {
	var sizeJobs []SizeJob
	var result []string

	var fileLen = len(j.Files)
	sizeJobs = make([]SizeJob, 0, fileLen)
	result = make([]string, 0, fileLen)

	for _, f := range j.Files {
		file, _ := os.Stat(f)
		sizeJobs = append(sizeJobs, SizeJob{f, file.Size()})
	}

	// sort by size: 0, 1, 2, 3 ...
	if !j.Reverse {
		sort.Slice(sizeJobs, func(i, j int) bool {
			return sizeJobs[i].Size < sizeJobs[j].Size
		})
	}
	// reverse sort by size: 100, 99, 98 ...
	if j.Reverse {
		sort.Slice(sizeJobs, func(i, j int) bool {
			return sizeJobs[i].Size > sizeJobs[j].Size
		})
	}

	for _, v := range sizeJobs {
		result = append(result, v.Name)
	}

	return result
}
