package ast

import (
	"fmt"
	"github.com/yinggaozhen/GoLearn/ast/controller"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// 参考资料 : http://www.01happy.com/golang-ast-use/

var hc controller.HelloController
func TestParseComment(t *testing.T) {
	parseFile := "/data/github/GoLearn/ast/controller/hello_comment.go"

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, parseFile, nil, parser.ParseComments)

	fmt.Println(f.Decls[2].(*ast.FuncDecl).Doc.List[0].Text)
	fmt.Println(f.Decls[2].(*ast.FuncDecl).Doc.List[1].Text)
	fmt.Println(f.Decls[2].(*ast.FuncDecl).Doc.List[2].Text)
}
