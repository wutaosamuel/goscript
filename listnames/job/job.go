package job

import (
	"goscript/listnames/config"
)

// Job is job
type Job struct {
	Files         []string // all files
	SelectedFiles []string // select files
	Reverse       bool     // reverse order
	IsSelect      bool     // check if it needs to select
}

// NewJob create Job object
func NewJob() *Job {
	return &Job{
		Files:         make([]string, 0),
		SelectedFiles: make([]string, 0),
		Reverse:       false,
		IsSelect:      false,
	}
}

// NewJobRead create job object by variables
func NewJobRead(files []string, pick []int, reverse bool) *Job {
	j := NewJob()
	j.Files = files
	j.Reverse = reverse
	if len(pick) != 0 {
		j.IsSelect = true
	}

	return j
}

// ReadConfig create job and read data from config
func ReadConfig(c *config.Config) *Job {
	j := NewJob()
	j.ReadConfig(c)

	return j
}

// ReadConfig read data from config
func (j *Job) ReadConfig(c *config.Config) {
	j.Files = c.Files
	j.Reverse = c.Reverse
	if len(c.Pick) != 0 {
		j.IsSelect = true
	}

	return
}

// SortDefault sort by default
func (j *Job) SortDefault() []string {
	if j.Reverse {
		fileLen := len(j.Files)
		files := make([]string, fileLen)
		for k, v := range j.Files {
			files[fileLen-k-1] = v
		}
		return files
	}
	return j.Files
}
