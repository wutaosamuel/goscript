package common

// version for now
var (
	version = "0.0.1"
)

// Version print out version
func Version() string {
	return version
}

// OpCode op code
type OpCode uint

// operation
const (
	DefaultOp OpCode = iota
	AddOp
	DeleteOp
	RenameOp
	ListOp
	CountOp
)
