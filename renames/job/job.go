package job

import (
	"errors"
	"io"
	"os"
	"sort"
	"sync"

	"github.com/wutaosamuel/goscript/renames/common"
	"github.com/wutaosamuel/goscript/renames/config"
)

// Job job object
type Job struct {
	Files         []string // all files
	SelectedFiles []string // select files by pick flag
	Pick          []int    // pick files
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
		Pick:          make([]int, 0),
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
	j.Pick = pick
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
	j.Pick = c.Pick
	j.AtBegin = c.Begin
	j.DoList = c.ListFile
	j.Reverse = c.Reverse

	doList := c.OpCode
	if doList == common.DefaultOp {
		j.ListOperation = c.List.Operation
	}
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

// SelectFiles select files
// 	- start from 0
func SelectFiles(files []string, pick []int) []string {
	var isStart bool
	var pLocation int
	var result []string
	// do not nothing, when no files or selection
	if len(files) == 0 || len(pick) == 0 {
		return files
	}

	isStart = false
	pLocation = 0
	sort.Ints(pick)
	for k, v := range files {
		if pLocation < len(pick) {
			if k == pick[pLocation] {
				isStart = !isStart
				pLocation++
			}
		}
		if isStart {
			result = append(result, v)
		}
	}

	return result
}

// GoCopy go copy
func GoCopy(src, dst string, wg *sync.WaitGroup) error {
	if err := Copy(src, dst); err != nil {
		return err
	}
	wg.Done()

	return nil
}

// Copy copy files
func Copy(src, dst string) error {
	// check file
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcStat.Mode().IsRegular() {
		return errors.New("CopyFile: " + src + " is non-regular source file ")
	}

	// copy contents
	return CopyContents(src, dst)
}

// CopyContents copy file content
func CopyContents(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return errors.New("CopyContents: open error" + "on " + src)
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}
