package job

import (
	"path/filepath"
	"strconv"
)

// Job object
type Job struct {
	Files []string // all files
	Char  string   // match char
}

// NewJob create job object
func NewJob() *Job {
	return &Job{
		Files: make([]string, 0),
		Char:  "",
	}
}

// SetJob set job
func (j *Job) SetJob(files []string, char string) {
	j.Files = files
	j.Char = char
}

// CountDefault count files' name by default
func (j *Job) CountDefault() []string {
	var result []string

	result = make([]string, 0, len(j.Files))
	for _, f := range j.Files {
		extension := filepath.Ext(f)
		base := filepath.Base(f)
		base = base[0 : len(base)-len(extension)]
		result = append(result, strconv.Itoa(len(base))+extension)
	}

	return result
}

// CountMatch count files' name by matching
func (j *Job) CountMatch() []string {
	var result []string

	if j.Char == "" || len(j.Files) == 0 {
		return result
	}

	for _, f := range j.Files {
		extension := filepath.Ext(f)
		base := filepath.Base(f)
		base = base[0 : len(base)-len(extension)]
		matcher := NewMatcher()

		matcher.Match(base, j.Char)
		result = append(result, matcher.String()+extension)
	}

	return result
}
