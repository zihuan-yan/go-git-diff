// Package parser 元数据
package parser

import "strings"

const (
	MetaIndex = "index" // 索引
)

// MetaType 元数据类型
type MetaType string

// Meta diff元数据
type Meta struct {
	MetaType MetaType `json:"meta_type,omitempty"` // 元数据类型
	Content  string   `json:"content,omitempty"`   // 内容
}

// NewMeta 创建diff元数据实例
func NewMeta() *Meta {
	return &Meta{}
}

// IsMeta 是否为diff元数据行
func IsMeta(diffLine string) bool {
	return strings.HasPrefix(diffLine, MetaIndex)
}

// Parse 解析diff元数据
func (m *Meta) Parse(diffLine string) *Meta {
	switch {
	case strings.HasPrefix(diffLine, MetaIndex):
		m.MetaType = MetaIndex
		m.Content = strings.TrimPrefix(diffLine, MetaIndex+Space)
	}

	return m
}
