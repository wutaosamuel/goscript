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

// ListOperationToInt list to int
var ListOperationToInt = map[ListOperation]int{
	DefaultList:   0,
	NameList:      1,
	TimeList:      2,
	SizeList:      3,
	ExtensionList: 4,
}

// IntToListOperation int to list
var IntToListOperation = map[int]ListOperation{
	0: DefaultList,
	1: NameList,
	2: TimeList,
	3: SizeList,
	4: ExtensionList,
}
