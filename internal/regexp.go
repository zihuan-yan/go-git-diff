// Package internal regexp.go
package internal

import "regexp"

// regexpT is a regexp template struct.
type regexpT struct {
	ChunkRegexp *regexp.Regexp
}

// RegexpT is a regexp template struct.
var RegexpT *regexpT

// InitRegexpT init regexp template.
func InitRegexpT() {
	RegexpT = &regexpT{}

	RegexpT.ChunkRegexp = regexp.MustCompile(`@@ -(\d+),?(\d+)? \+(\d+),?(\d+)? @@ ?(.+)?`)
}
