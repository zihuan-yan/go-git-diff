// Package parser 图例
package parser

import "strings"

const (
	MarkDelete = "---" // 源文件图例
	MarkInsert = "+++" // 目标文件图例
)

// Legend diff图例
type Legend struct {
	Mark    string `json:"mark,omitempty"`
	Content string `json:"content,omitempty"`
}

// NewLegend 创建diff图例实例
func NewLegend() *Legend {
	return &Legend{}
}

// IsLegend 是否为diff图例起始行
func IsLegend(diffLine string) bool {
	return strings.HasPrefix(diffLine, MarkDelete) || strings.HasPrefix(diffLine, MarkInsert)
}

// Parse 解析diff图例
func (l *Legend) Parse(diffLine string) *Legend {
	if len(diffLine) < 4 {
		return nil
	}

	l.Mark = diffLine[:3]
	l.Content = diffLine[4:]

	return l
}
