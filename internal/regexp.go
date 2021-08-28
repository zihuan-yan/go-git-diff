// Package internal 正则表达式
package internal

import "regexp"

type regexpT struct {
	ChunkRegexp *regexp.Regexp
}

// RegexpT 正则表达式模版
var RegexpT *regexpT

// InitRegexpT 初始化正则表达式模板
func InitRegexpT() {
	RegexpT = &regexpT{}

	RegexpT.ChunkRegexp = regexp.MustCompile(`@@ -(\d+),?(\d+)? \+(\d+),?(\d+)? @@ ?(.+)?`)
}
