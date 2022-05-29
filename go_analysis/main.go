package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

// перевести на мап[тип]индекс
var correctOrder = []string{
	"Document",
	"Array",
	"float64",
	"string",
	"types.Binary",
	"types.ObjectID",
	"bool",
	"time.Time",
	"types.NullType",
	"types.Regex",
	"int32",
	"types.Timestamp",
	"int64",
}

func main() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".go" {
				find(path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

}

func indexSlice(typeCase string) int {
	for idx, tp := range correctOrder {
		if tp == typeCase {
			return idx
		}
	}

	return -1
}

func find(path string) {

	var idx int

	//	goFile := "test_data/example.go"

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// ast.Print(fset, f)

	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.CommentGroup:
		// 	fmt.Printf("%+v\n", n)
		case *ast.TypeSwitchStmt:
			var name string

			for _, el := range n.Body.List {
				for _, cs := range el.(*ast.CaseClause).List {
					switch cs := cs.(type) {
					case *ast.StarExpr:
						if sexp, ok := cs.X.(*ast.SelectorExpr); ok {
							name = sexp.Sel.Name
							// name = fmt.Sprintf("%s.%s", sexp.X.(*ast.Ident).Name, sexp.X.(*ast.Ident).Name)
						}

					case *ast.SelectorExpr:
						name = fmt.Sprintf("%s.%s", cs.X, cs.Sel.Name)

					case *ast.Ident:
						name = cs.Name
					}

					iSl := indexSlice(name)
					if iSl < idx {
						//fmt.Printf("%+v\n", n.Body.List)
						fmt.Println("неправильный порядок switch:", fset.Position(n.Switch))
					}
					idx = iSl

				}
			}
		}

		return true
	})
}
