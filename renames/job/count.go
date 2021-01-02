package job

import (
	"../../countnames"
)

// Count count file names
func (j *Job) Count() {
	// list file order first
	j.list()

	if len(j.SelectedFiles) != 0 {
		result := countnames.Execute(j.SelectedFiles, j.Char)
		countnames.ConsoleString(result, j.SelectedFiles)
	}
	if len(j.SelectedFiles) == 0 {
		result := countnames.Execute(j.Files, j.Char)
		countnames.ConsoleString(result, j.Files)
	}
}
