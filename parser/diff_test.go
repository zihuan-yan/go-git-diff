// Package parser 模块说明
package parser

import (
	"encoding/json"
	"testing"
)

var diffText = `diff --git a/main/tokenAnalyze.go b/main/tokenAnalyze.go
index 9eec139..ac55e8d 100644
--- a/main/tokenAnalyze.go
+++ b/main/tokenAnalyze.go
@@ -12,7 +12,7 @@ import (
 )

 type myVisitor struct {
-       fSet *token.FileSet
+       fSet *token.FileSet 1
 }

 func (v *myVisitor) Visit(n ast.Node) (w ast.Visitor) {
@@ -48,16 +48,22 @@ func getFileChange(diffStr string) *model.FileChange {
 func main() {

        out, _, _ := diff()
+
+       s := strings.Split(out, "\n+++ ")
+       for _, item := range s {
+               fmt.Println("11111111111111111111")
+               fmt.Println(item)
+               fmt.Println("2222222222222222222")
+       }
+
        // projectInfo := make([]model.ProjectInfo, 0)

-       firstIndex := strings.Index(out, "+++")
-       out = out[firstIndex:]
-       firstIndex1 := strings.Index(out, " ")
-       firstIndex2 := strings.Index(out, "\n")
-       out = out[firstIndex1+1 : firstIndex2]
-       fmt.Println(out)
-       s := "abc"
-       fmt.Println(strings.Index(s, "abc"))
+       // firstIndex := strings.Index(out, "+++")
+       // out = out[firstIndex:]
+       // firstIndex1 := strings.Index(out, " ")
+       // firstIndex2 := strings.Index(out, "\n")
+       // out = out[firstIndex1+1 : firstIndex2]
+       // fmt.Println(out)
        // fset := token.NewFileSet()
        // f, _ := parser.ParseFile(fset, "inter.go", nil, parser.AllErrors)
        // visitor := new(myVisitor)
diff --git a/model/test.go b/model/test.go
index 9ee951e..5eaae2f 100644
--- a/model/test.go
+++ b/model/test.go
@@ -1,3 +1,6 @@
 package model

 func test() {}
+adkn
+
+lsndlsdf`

func Test_diff_Parse(t *testing.T) {
	t.Run("test diff", func(t *testing.T) {
		testDiff := NewDiff()
		testDiff.Parse(diffText)
		marshal, err := json.MarshalIndent(testDiff, "", "  ")
		if err != nil {
			t.Error(err.Error())
			return
		}

		t.Logf("diff:\n%+v", string(marshal))
	})
}
