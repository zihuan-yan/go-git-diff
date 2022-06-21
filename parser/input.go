// Package parser 输入数据源
package parser

import "strings"

const (
	InputPrefix = "diff --git " // diff输入前缀
)

// Input diff输入
type Input struct {
	Source string `json:"source,omitempty"` // 源文件
	Target string `json:"target,omitempty"` // 目标文件
}

// NewInput 创建diff输入实例
func NewInput() *Input {
	return &Input{}
}

// IsInput 是否为diff输入行
func IsInput(diffLine string) bool {
	return strings.HasPrefix(diffLine, InputPrefix)
}

// Parse 解析diff输入
func (i *Input) Parse(diffLine string) *Input {
	inputComment := strings.TrimPrefix(diffLine, InputPrefix)
	inputComments := strings.Split(inputComment, Space)
	if len(inputComments) < 2 {
		return nil
	}

	i.Source = inputComments[0]
	i.Target = inputComments[1]

	return i
}
