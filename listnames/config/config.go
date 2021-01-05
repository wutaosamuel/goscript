package config

// Config contain filenames and operations
type Config struct {
	Files     []string // The input files' name
	Reverse   bool     // Reverse orders
	Pick      []int    // Select files in the range

	Operation OpCode // Operation code for subcommand

	Result []string // the final result, which is output files' name
}

// NewConfig create a new Config object
func NewConfig() *Config {
	return &Config{
		Files:     make([]string, 0),
		Reverse:   false,
		Pick:      make([]int, 0),
		Operation: DefaultOp,
		Result:    make([]string, 0),
	}
}
