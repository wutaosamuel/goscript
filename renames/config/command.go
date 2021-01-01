package config

// Add add command object
type Add struct {
	Char   string // Characters for adding
	Number int    // Number for adding
}

// Delete delete command object
type Delete struct {
	Number int // Number for deleting filenames
}

// Rename rename command object
type Rename struct {
	Char      string // Characters to change
	Number    int    // Number at filename
	Extension string // change file extension
}

// List list command object
type List struct {
	Operation ListOperation // Operation
}

// Count count command object
type Count struct {
	Char string
}

// NewAdd create Add object
func NewAdd() *Add {
	return &Add{
		Char:   "",
		Number: 0,
	}
}

// NewDelete create Delete object
func NewDelete() *Delete {
	return &Delete{
		Number: 0,
	}
}

// NewRename create Rename object
func NewRename() *Rename {
	return &Rename{
		Char:      "",
		Number:    0,
		Extension: "",
	}
}

// NewList create List object
func NewList() *List {
	return &List{
		Operation: DefaultList,
	}
}

// NewCount create Count object
func NewCount() *Count {
	return &Count{
		Char: "",
	}
}

// SetListOperation set list operation
func (l *List) SetListOperation(operation int) {
	l.Operation = setListOperation(operation)
}
