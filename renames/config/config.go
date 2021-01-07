package config

import "github.com/wutaosamuel/goscript/renames/common"

// Config contain filenames and operations
type Config struct {
	Files     []string      // The input files' name
	OutputDir string        // Output directory for copying
	Begin     bool          // Do operation at begin of filename
	Pick      []int         // Select files in the range
	ListFile  ListOperation // List Operation
	Reverse   bool          // Reverse orders

	OpCode common.OpCode // opcode for sub command

	// subcommand
	Add    *Add
	Delete *Delete
	Rename *Rename
	List   *List
	Count  *Count

	Result []string // the final result, which is output files' name
}

// NewConfig create a new Config object
func NewConfig() *Config {
	return &Config{
		Files:     make([]string, 0),
		OutputDir: "",
		Begin:     false,
		Pick:      make([]int, 0),
		ListFile:  DefaultList,
		Reverse:   false,

		OpCode: common.DefaultOp,

		Add:    NewAdd(),
		Delete: NewDelete(),
		Rename: NewRename(),
		List:   NewList(),
		Count:  NewCount(),

		Result: make([]string, 0),
	}
}

// SetListOperation set list operation
func (c *Config) SetListOperation(operation int) {
	c.ListFile = setListOperation(operation)
}

// GetListFileInt get list operation
func (c *Config) GetListFileInt() int {
	return ListOperationToInt[c.ListFile]
}

// setListOperation set list operation
func setListOperation(operation int) ListOperation {
	if operation < 0 || operation > 4 {
		return DefaultList
	}
	return IntToListOperation[operation]
}

// ListOperationInt get list operation int
func ListOperationInt(operation ListOperation) int {
	return ListOperationToInt[operation]
}
