package job

import (
	"path/filepath"
)

// FileJob slice of files
type FileJob struct {
	Files []*FileName // all file name
}

// NewFileJob create new FileJob
func NewFileJob() *FileJob {
	return &FileJob{
		Files: make([]*FileName, 0),
	}
}

// ParseFilenames parse filename
func (j *FileJob) ParseFilenames(files []string) {
	for k, v := range files {
		filename := NewFileName()
		dir := filepath.Dir(v)
		base := filepath.Base(v)
		extension := filepath.Ext(base)

		filename.ID = k
		filename.Dir = dir
		filename.Name = base[0 : len(base)-len(extension)]
		filename.Extension = extension
	}
}

// ToFileNames to filenames
// 	- files, contain
// TODO: check local dir may contain the same name file
func (j *FileJob) ToFileNames(dir string) ([]string, bool) {
	var result []string
	// do not need to copy
	if dir == "" {
		for _, v := range j.Files {
			result = append(result, v.ToFullFileName())
		}
		return result, false
	}

	// copy to target dir, check the filenames
	filenames := make(map[string]bool, 0)
	for _, v := range j.Files {
		filename := v.ToFileName()
		if filenames[filename] {
			return []string{}, true
		}
		filenames[filename] = true
		result = append(result, filepath.Join(dir, filename))
	}

	return result, false
}

// FileName including parent dir, base name and extension
type FileName struct {
	ID        int    // file order
	Dir       string // parent directory path
	Name      string // base name without extension
	Extension string // extension
}

// NewFileName create new FileName object
func NewFileName() *FileName {
	return &FileName{
		ID:        -1,
		Dir:       "",
		Name:      "",
		Extension: "",
	}
}

// ToFileName only filename without dir path
func (f *FileName) ToFileName() string {
	return f.Name + f.Extension
}

// ToFullFileName filename with dir
func (f *FileName) ToFullFileName() string {
	return filepath.Join(f.Dir, f.ToFileName())
}
