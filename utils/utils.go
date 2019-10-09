package utils

import (
	"errors"
	"strings"
)

var (
	// OrchThriftPrefix is the prefix to identify if a thrift path is the thrift content
	OrchThriftPrefix = "THRIFT@"
)

// IsFileContent identifies if f contains 'THRIFT@'
func IsFileContent(f string) bool {
	if f != "" && strings.Contains(f, OrchThriftPrefix) {
		return true
	}
	return false
}

// RemoveOrchThriftPrefix removes the co-orch-thrift prefix and everything before it
func RemoveOrchThriftPrefix(f string) ([]byte, error) {
	slice := strings.Split(f, OrchThriftPrefix)
	if len(slice) != 2 {
		return nil, errors.New("Error removing co-orch-thrift prefix")
	}
	return []byte(slice[1]), nil
}
