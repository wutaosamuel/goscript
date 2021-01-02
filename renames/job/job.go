package job

import (
	"../common"
	"../config"
)

// Job job object
type Job struct {
	Files         []string // all files
	SelectedFiles []string // select files by pick flag
	IsSelect      bool     // check if it needs to select
	OutputDir     string   // output dir
	AtBegin       bool     // do operation on start of filename

	DoList  config.ListOperation // use listed files order:default, add, delete, rename, list, count action
	Reverse bool                 // reverse order

	Char          string               // Characters
	Number        int                  // Number of characters or Integer number
	Extension     string               // file extension
	ListOperation config.ListOperation // list operation
}

// NewJob create job object
func NewJob() *Job {
	return &Job{
		Files:         make([]string, 0),
		SelectedFiles: make([]string, 0),
		IsSelect:      false,
		OutputDir:     "",
		AtBegin:       false,

		DoList:  config.DefaultList,
		Reverse: false,

		Char:          "",
		Number:        -1,
		Extension:     "",
		ListOperation: config.DefaultList,
	}
}

// NewJobRead create job object by variables
func NewJobRead(files []string, outputDir string, pick []int, atBegin bool, doList config.ListOperation,
	reverse bool, char string, number int, extension string, listOperation config.ListOperation) *Job {
	j := NewJob()
	j.Files = files
	j.OutputDir = outputDir
	if len(pick) != 0 {
		j.IsSelect = true
	}
	j.AtBegin = atBegin
	j.DoList = doList
	j.Reverse = reverse
	j.Char = char
	j.Number = number
	j.Extension = extension
	j.ListOperation = listOperation

	return j
}

// ReadConfig create job from config
func ReadConfig(c *config.Config) *Job {
	j := NewJob()
	j.ReadConfig(c)

	return j
}

// ReadConfig read from config
func (j *Job) ReadConfig(c *config.Config) {
	j.Files = c.Files
	j.OutputDir = c.OutputDir
	if len(c.Pick) != 0 {
		j.IsSelect = true
	}
	j.AtBegin = c.Begin
	j.DoList = c.ListFile
	j.Reverse = c.Reverse

	doList := common.DefaultOp
	if doList == common.AddOp {
		j.Char = c.Add.Char
		j.Number = c.Add.Number
	}
	if doList == common.DeleteOp {
		j.Number = c.Delete.Number
	}
	if doList == common.RenameOp {
		j.Char = c.Rename.Char
		j.Number = c.Rename.Number
		j.Extension = c.Rename.Extension
	}
	if doList == common.ListOp {
		j.ListOperation = c.List.Operation
	}
	if doList == common.CountOp {
		j.Char = c.Count.Char
	}
}
