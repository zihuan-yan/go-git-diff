// Package parser 处理go语言git diff解析 `git diff source target`
package parser

import (
	"strings"

	"github.com/lower-coder/go-git-diff/internal"
)

const (
	FileDeleted  FileMode = 1 // 删除
	FileModified FileMode = 2 // 修改
	FileInserted FileMode = 3 // 新增
)

// FileMode 文件变动模式
type FileMode int

// File 文件
type File struct {
	Mode    FileMode  `json:"mode,omitempty"`    // 文件变动模式
	Input   *Input    `json:"input,omitempty"`   // diff输入
	Metas   []*Meta   `json:"metas,omitempty"`   // diff元数据
	Legends []*Legend `json:"legends,omitempty"` // diff图例
	Chunks  []*Chunk  `json:"chunks,omitempty"`  // diff块
}

// NewFile 创建文件实例
func NewFile() *File {
	return &File{
		Metas:   []*Meta{},
		Chunks:  []*Chunk{},
		Legends: []*Legend{},
	}
}

// Diff diff
type Diff struct {
	Files []*File `json:"files,omitempty"` // 文件列表
}

// NewDiff 创建diff实例
func NewDiff() *Diff {
	internal.InitRegexpT() // 初始化正则表达式模板

	return &Diff{
		Files: []*File{},
	}
}

// Parse 解析diff
func (d *Diff) Parse(diffText string) *Diff {
	// 按行解析
	diffLines := strings.Split(diffText, "\n")
	diffFile := NewFile()
	diffChunk := NewChunk()

	var insertNO int32 // 新增行号
	var deleteNO int32 // 删除行号
	var inChunk bool   // 是否在diff块里

	for _, diffLine := range diffLines {
		switch {
		case IsInput(diffLine):
			diffInput := NewInput()
			diffInput.Parse(diffLine)

			diffFile = NewFile()
			diffFile.Input = diffInput
			diffFile.Mode = FileModified

			inChunk = false // 进入diff块之后设置为true

			d.Files = append(d.Files, diffFile)
		case IsMeta(diffLine):
			diffFile.Metas = append(diffFile.Metas, NewMeta().Parse(diffLine))
		case diffLine == "+++ /dev/null":
			diffFile.Mode = FileDeleted
		case diffLine == "--- /dev/null":
			diffFile.Mode = FileInserted
		case IsLegend(diffLine):
			diffFile.Legends = append(diffFile.Legends, NewLegend().Parse(diffLine))
		case IsChunk(diffLine):
			inChunk = true
			diffChunk = NewChunk()
			diffChunk.Parse(diffLine)

			diffFile.Chunks = append(diffFile.Chunks, diffChunk)

			insertNO = diffChunk.TargetRange.Start // 删除的代码在源文件里
			deleteNO = diffChunk.SourceRange.Start // 新增的代码在目标文件里
		case inChunk:
			line := Line{}
			line.LineMode(diffLine)

			sourceLine := line
			targetLine := line

			// 删除的代码在源文件里、新增的代码在目标文件里
			switch line.Mode {
			case LineInsert:
				targetLine.Number = insertNO
				diffChunk.TargetRange.Lines = append(diffChunk.TargetRange.Lines, &targetLine)
				insertNO++
			case LineDelete:
				sourceLine.Number = deleteNO
				diffChunk.SourceRange.Lines = append(diffChunk.SourceRange.Lines, &sourceLine)
				deleteNO++
			default:
				targetLine.Number = insertNO
				diffChunk.TargetRange.Lines = append(diffChunk.TargetRange.Lines, &targetLine)
				insertNO++

				sourceLine.Number = deleteNO
				diffChunk.SourceRange.Lines = append(diffChunk.SourceRange.Lines, &sourceLine)
				deleteNO++
			}
		}
	}

	return d
}
