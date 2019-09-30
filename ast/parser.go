package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// 参考资料 : http://www.01happy.com/golang-ast-use/
// 获取注释
func main() {
	parseFile := "/data/github/GoLearn/ast/controller/hello_comment.go"

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, parseFile, nil, parser.ParseComments)

	fmt.Println(f.Decls[2].(*ast.FuncDecl).Doc.List[0].Text)
	fmt.Println(f.Decls[2].(*ast.FuncDecl).Doc.List[1].Text)
	fmt.Println(f.Decls[2].(*ast.FuncDecl).Doc.List[2].Text)
}