// Package diff diff块
package diff

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/lower-coder/go-git-diff/internal"
)

const (
	ChunkPrefix = "@@" // diff块前缀
)

// Chunk diff块
type Chunk struct {
	Header      string `json:"header,omitempty"`       // 块头部
	SourceRange *Range `json:"source_range,omitempty"` // 源文件行范围
	TargetRange *Range `json:"target_range,omitempty"` // 目标文件行范围
}

// Range 行范围
type Range struct {
	Start  int32   `json:"start,omitempty"`  // 起始行号
	Length int32   `json:"length,omitempty"` // 范围长度
	Lines  []*Line `json:"lines,omitempty"`  // 行列表
}

// NewChunk 创建diff块实例
func NewChunk() *Chunk {
	return &Chunk{}
}

// IsChunk 是否为diff块起始行
func IsChunk(diffLine string) bool {
	return strings.HasPrefix(diffLine, ChunkPrefix)
}

// Parse 解析diff块
func (c *Chunk) Parse(diffLine string) *Chunk {
	chunkRegexpGroup := internal.RegexpT.ChunkRegexp.FindStringSubmatch(diffLine)
	if len(chunkRegexpGroup) < 6 {
		return nil
	}

	c.Header = chunkRegexpGroup[5]
	c.SourceRange = &Range{
		Start:  cast.ToInt32(chunkRegexpGroup[1]),
		Length: cast.ToInt32(chunkRegexpGroup[2]),
		Lines:  []*Line{},
	}
	c.TargetRange = &Range{
		Start:  cast.ToInt32(chunkRegexpGroup[3]),
		Length: cast.ToInt32(chunkRegexpGroup[4]),
		Lines:  []*Line{},
	}

	return c
}
