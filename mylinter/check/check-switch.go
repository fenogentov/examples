package main

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// go install check/check-switch.go
// go vet -vettool=$(which check-switch) ./...

// orderTypes is preferred order of the types in the switch.
var orderTypes = map[string]int{
	"Document":  0,
	"Array":     1,
	"float64":   2,
	"string":    3,
	"Binary":    4,
	"ObjectID":  5,
	"bool":      6,
	"time.Time": 7,
	"NullType":  8,
	"Regex":     9,
	"int32":     10,
	"Timestamp": 11,
	"int64":     12,
}

var Analyzer = &analysis.Analyzer{
	Name: "checkerswitch",
	Doc:  "checking the preferred order of types in the switch",
	Run:  run,
}

func main() {
	singlechecker.Main(Analyzer)
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			var idx int
			switch n := n.(type) {
			case *ast.CommentGroup:
			//	fmt.Printf("%+v\n", n)
			case *ast.TypeSwitchStmt:
				var name string
				for _, el := range n.Body.List {
					if len(el.(*ast.CaseClause).List) < 1 {
						continue
					}
					firstTypeCase := el.(*ast.CaseClause).List[0]
					switch firstTypeCase := firstTypeCase.(type) {
					case *ast.StarExpr:
						if sexp, ok := firstTypeCase.X.(*ast.SelectorExpr); ok {
							name = sexp.Sel.Name
							// name = fmt.Sprintf("%s.%s", sexp.X.(*ast.Ident).Name, sexp.X.(*ast.Ident).Name)
						}
					case *ast.SelectorExpr:
						name = fmt.Sprintf("%s.%s", firstTypeCase.X, firstTypeCase.Sel.Name)

					case *ast.Ident:
						name = firstTypeCase.Name
					}

					idxSl, ok := orderTypes[name]
					if ok && (idxSl < idx) {
						pass.Reportf(n.Pos(), "non-observance of the preferred order of types")
					}
					idx = idxSl

					if len(el.(*ast.CaseClause).List) > 1 {
						subidx := idx
						for i := 0; i < len(el.(*ast.CaseClause).List); i++ {
							cs := el.(*ast.CaseClause).List[i]
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

							subidxSl, ok := orderTypes[name]
							if ok && (subidxSl < subidx) {
								pass.Reportf(n.Pos(), "non-observance of the preferred order of types")
							}
							subidx = subidxSl
						}
					}
				}
			}

			return true
		})
	}

	return nil, nil
}
