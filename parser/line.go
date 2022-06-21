// Package parser diff行
package parser

const (
	InsertPrefix = "+" // 新增行前缀
	DeletePrefix = "-" // 删除行前缀

	LineUnchanged LineMode = 1 // 上下文
	LineInsert    LineMode = 2 // 新增
	LineDelete    LineMode = 3 // 删除
)

// LineMode 行变动类型
type LineMode int

// Line diff行
type Line struct {
	Mode    LineMode `json:"mode,omitempty"`    // 行变动类型
	Number  int32    `json:"number,omitempty"`  // 行号
	Content string   `json:"content,omitempty"` // 内容
}

// LineMode 设置行变动类型
func (l *Line) LineMode(diffLine string) {
	if len(diffLine) < 1 {
		l.Mode = LineUnchanged
		return
	}

	l.Content = diffLine[1:]

	switch diffLine[:1] {
	case InsertPrefix:
		l.Mode = LineInsert
	case DeletePrefix:
		l.Mode = LineDelete
	default:
		l.Mode = LineUnchanged
	}
}
