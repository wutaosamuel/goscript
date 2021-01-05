package job

import (
	"path/filepath"
	"sort"
)

// ExtensionJob extension job object
type ExtensionJob struct {
	Extension string   // file extension
	BaseNames []string // files' base name
	IDs []int // files id
}

// SortExtension sort file by extension
func (j *Job) SortExtension() []string {
	var extensionJob []*ExtensionJob
	var result []string

	var fileLen = len(j.Files)
	extensionJob = make([]*ExtensionJob, 0)
	result = make([]string, 0, fileLen)

	for k, v := range j.Files {
		extension := filepath.Ext(v)
		base := v[0 : len(v)-len(extension)]
		isNew := true

		for _, e := range extensionJob {
			if extension == e.Extension {
				isNew = false
				e.BaseNames = append(e.BaseNames, base)
				e.IDs = append(e.IDs, k)
				break
			}
		}

		if isNew {
			extensionJob = append(extensionJob, &ExtensionJob{extension, []string{base}, []int{k}})
		}
	}

	// Sort files' name in same extension
	for _, e := range extensionJob {
		sort.Strings(e.BaseNames)
	}
	// Sort by extension, a, b, c ...
	if !j.Reverse {
		sort.Slice(extensionJob, func(i, j int) bool {
			return extensionJob[i].Extension < extensionJob[j].Extension
		})
	}
	// Sort by extension, z, y, x ...
	if j.Reverse {
		sort.Slice(extensionJob, func(i, j int) bool {
			return extensionJob[i].Extension > extensionJob[j].Extension
		})
	}

	// result
	for _, e := range extensionJob {
		for _, v := range e.IDs {
			result = append(result, j.Files[v])
		}
	}

	return result
}
