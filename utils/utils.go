package utils

import "strings"

var (
	// OrchThriftPrefix is the prefix to identify if a thrift path is the thrift content
	OrchThriftPrefix = "thrift://"
)

// IsFileContent identifies if f starts with 'thrift://'
func IsFileContent(f string) bool {
	if f != "" && strings.HasPrefix(f, OrchThriftPrefix) {
		return true
	}
	return false
}
