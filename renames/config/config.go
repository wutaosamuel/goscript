package config

// Config contain filenames and operations
type Config struct {
	Files     []string      // The input files' name
	OutputDir string        // Output directory for copying
	Begin     bool          // Do operation at begin of filename
	Pick      []int         // Select files in the range
	Reverse   bool          // Reverse orders
	ListFile  ListOperation // List Operation

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
		Reverse:   false,
		ListFile:  DefaultList,

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

// setListOperation set list operation
func setListOperation(operation int) ListOperation {
	switch {
	case operation == 1:
		return NameList
	case operation == 2:
		return TimeList
	case operation == 3:
		return SizeList
	case operation == 4:
		return ExtensionList
	default:
		return DefaultList
	}
}
