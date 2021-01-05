package job

import (
	"os"
	"sort"
	"time"
)

// TimeJob time job
// TODO: need to support
// changed time
// modify time  x
// access time
type TimeJob struct {
	Name string    // file name
	Time time.Time // time
}

// SortMTime sort by modify time
func (j *Job) SortMTime() []string {
	var timeJobs []TimeJob
	var result []string

	var filesLen = len(j.Files)
	timeJobs = make([]TimeJob, 0, filesLen)
	result = make([]string, 0, filesLen)

	for _, f := range j.Files {
		file, _ := os.Stat(f)
		timeJobs = append(timeJobs, TimeJob{f, file.ModTime()})
	}

	// sort by modified time, 2000-01-01, 3000-01-01, 4000-01-01
	if !j.Reverse {
		sort.Slice(timeJobs, func(i, j int) bool {
			return timeJobs[i].Time.Before(timeJobs[j].Time)
		})
	}
	// sort by modified time, 4000-01-01, 3000-01-01, 2000-01-01
	if j.Reverse {
		sort.Slice(timeJobs, func(i, j int) bool {
			return timeJobs[i].Time.After(timeJobs[j].Time)
		})
	}

	for _, j := range timeJobs {
		result = append(result, j.Name)
	}

	return result
}
