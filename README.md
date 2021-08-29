# go-git-diff [![CI](https://github.com/lower-coder/go-git-diff/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/lower-coder/go-git-diff/actions/workflows/go.yml) [![CI](https://github.com/lower-coder/go-git-diff/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/lower-coder/go-git-diff/actions/workflows/codeql-analysis.yml)

A tool for developer to parse output of git diff command.

## Usage

- diff: A parser that parse output of git diff command to struct diff.

## Example

```go
package amin

import (
	"github.com/lower-coder/go-git-diff/diff"
)

const diffText = `diff --git a/file1 b/file1
index 504d2a1..50ccec3 100644
--- a/file1
+++ b/file1
@@ -1,4 +1,4 @@
+add a line
 some
 lines
-in
 file1`

func main() {
	d := diff.NewDiff()
	d.Parse(diffText)
}
```

[![Go Report Card](https://goreportcard.com/badge/github.com/lower-coder/go-git-diff)](https://goreportcard.com/report/github.com/lower-coder/go-git-diff)
[![license](https://img.shields.io/github/license/lower-coder/go-git-diff)](https://img.shields.io/github/license/lower-coder/go-git-diff)
[![go-version](https://img.shields.io/github/go-mod/go-version/lower-coder/go-git-diff)](https://img.shields.io/github/go-mod/go-version/lower-coder/go-git-diff)
[![contributors](https://img.shields.io/github/contributors/lower-coder/go-git-diff)](https://img.shields.io/github/contributors/lower-coder/go-git-diff)
[![stars](https://img.shields.io/github/stars/lower-coder/go-git-diff)](https://img.shields.io/github/stars/lower-coder/go-git-diff)
[![lines](https://img.shields.io/tokei/lines/github/lower-coder/go-git-diff)](https://img.shields.io/tokei/lines/github/lower-coder/go-git-diff)
[![downloads](https://img.shields.io/github/downloads/lower-coder/go-git-diff/total)](https://img.shields.io/github/downloads/lower-coder/go-git-diff/total)
