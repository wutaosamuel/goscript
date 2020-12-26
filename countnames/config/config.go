package config

// Config contain filename and operations
type Config struct {
	Files []string // The input files' name
	Char  string   // Characters to match
	Match bool     // Match or unMatch
}

// NewConfig create a new Config object
func NewConfig() *Config {
	return &Config{
		Files: make([]string, 0),
		Char:  "",
		Match: false,
	}
}
