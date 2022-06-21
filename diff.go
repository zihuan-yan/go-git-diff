// package diff diff.go.
package diff

import (
	"log"

	"github.com/lower-coder/go-git-diff/internal"
	"github.com/lower-coder/go-git-diff/parser"
)

// GitDiffStaged parse the output of `git diff --staged`.
//  @return *parser.Diff
//  @return error
func GitDiffStaged() (*parser.Diff, error) {
	cmd := internal.NewCommand("git", "diff", "--staged")
	output, err := cmd.Run()
	if err != nil {
		log.Printf("get diff staged failed, err = %+v", err)
		return nil, err
	}
	diff := parser.NewDiff()
	diff.Parse(output)
	return diff, nil
}
