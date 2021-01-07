package job

import "goscript/countnames"

// Count count file names
func (j *Job) Count() {
	// list file order first
	j.list(j.DoList)

	if len(j.SelectedFiles) != 0 {
		result := countnames.Execute(j.SelectedFiles, j.Char)
		countnames.ConsoleString(j.SelectedFiles, result)
	}
	if len(j.SelectedFiles) == 0 {
		result := countnames.Execute(j.Files, j.Char)
		countnames.ConsoleString(j.Files, result)
	}
}
