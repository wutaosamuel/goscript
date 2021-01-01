package config

// ListOperation list operation
type ListOperation uint

// DefaultOP
const (
	DefaultList ListOperation = iota
	NameList
	TimeList
	SizeList
	ExtensionList
)

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
