package job

import (
	"strconv"
	"strings"
)

// Matcher matcher object
type Matcher struct {
	AtStart       bool     // match at start string
	UnmatchFields []string // unmatched string
	AtEnd         bool     // match at end
	Pattern       string   // string to match
}

// NewMatcher create new Matcher object
func NewMatcher() *Matcher {
	return &Matcher{
		AtStart:       false,
		UnmatchFields: make([]string, 0),
		AtEnd:         false,
		Pattern:       "",
	}
}

// SetMatcher read from string
// TODO: not support like _abab_, which pattern is ab
func (m *Matcher) SetMatcher(message string, pattern string) {
	if message == "" {
		return
	}
	if pattern == "" || !strings.Contains(message, pattern) {
		m.Pattern = pattern
		m.UnmatchFields = append(m.UnmatchFields, message)
		return
	}

	// check start and end of string
	pLen := len(pattern)
	if message[:pLen] == pattern {
		m.AtStart = true
	}
	if message[len(message)-pLen:] == pattern {
		m.AtEnd = true
	}

	// get unMatched substring
	m.Pattern = pattern

	for _, v := range strings.Split(message, pattern) {
		if v == "" || v == " " {
			continue
		}
		m.UnmatchFields = append(m.UnmatchFields, v)
	}
}

// Match start to match
func (m *Matcher) Match(message string, pattern string) {
	m.SetMatcher(message, pattern)
}

// String string
func (m *Matcher) String() string {
	var result string

	// message == pattern
	if m.AtStart && m.AtEnd && len(m.UnmatchFields) == 0 {
		return m.Pattern
	}
	// message not contain pattern
	if !m.AtStart && !m.AtEnd && len(m.UnmatchFields) == 1 {
		return strconv.Itoa(len(m.UnmatchFields[0]))
	}

	// pattern at start
	if m.AtStart {
		result = m.Pattern + "_"
	}
	// pattern at mid
	for _, v := range m.UnmatchFields {
		n := strconv.Itoa(len(v))
		result += n + "_" + m.Pattern + "_"
	}
	// pattern at end
	if !m.AtEnd {
		result = result[:len(result)-len(m.Pattern)-2]
	}
	if m.AtEnd {
		result = result[:len(result)-1]
	}

	return result
}
