package config

// OpCode operation code for subcommand
type OpCode uint

// Subcommand operations
const (
	ErrorOp OpCode = iota
	DefaultOp
	NameOp
	TimeOp
	SizeOp
	ExtensionOp
)
