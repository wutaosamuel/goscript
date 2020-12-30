package job

import "testing"

// TestSetMatcher tet set matcher
func TestSetMatcher(t *testing.T) {
	var pattern string = "test"
	var message []string = []string {
		"a",
		"atest",
		"file",
		"test",
		"filetest",
		"filetestfile",
		"testfile",
		"testfiletest",
		"filetesttestfile", // not support yet
		"testfiletestfile",
	}
	
	for _, v := range message {
		m := NewMatcher()
		m.SetMatcher(v, pattern)
		t.Log(v, " -> ", m.String())
	}
}