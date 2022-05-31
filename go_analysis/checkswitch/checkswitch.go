package checkswitch

import (
	"flag"
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:             "checkswitch",
	Doc:              "reports checkswitch",
	Flags:            flag.FlagSet{},
	Run:              run,
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{},
	ResultType:       nil,
	FactTypes:        []analysis.Fact{},
}

// перевести на мап[тип]индекс
var correctOrder = []string{
	"Document",
	"Array",
	"float64",
	"string",
	"Binary",
	"ObjectID",
	"bool",
	"time.Time",
	"NullType",
	"Regex",
	"int32",
	"Timestamp",
	"int64",
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			var idx int
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
						if iSl == -1 {
							continue
						}
						if iSl < idx {
							fmt.Println("неправильный порядок switch:") //, fset.Position(n.Switch), name)
						}
						idx = iSl

					}
				}
			}

			return true
		})

	}

	return nil, nil
}

// func main() {
// 	err := filepath.Walk(".",
// 		func(path string, info os.FileInfo, err error) error {
// 			if err != nil {
// 				return err
// 			}
// 			if !info.IsDir() && filepath.Ext(path) == ".go" {
// 				find(path)
// 			}
// 			return nil
// 		})
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func indexSlice(typeCase string) int {
	for idx, tp := range correctOrder {
		if tp == typeCase {
			return idx
		}
	}

	return -1
}

// func find(path string) {

// 	var idx int

// 	//	goFile := "test_data/example.go"

// 	fset := token.NewFileSet()
// 	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// ast.Print(fset, f)

// 	ast.Inspect(f, func(n ast.Node) bool {
// 		switch n := n.(type) {
// 		case *ast.CommentGroup:
// 		// 	fmt.Printf("%+v\n", n)
// 		case *ast.TypeSwitchStmt:
// 			var name string

// 			for _, el := range n.Body.List {
// 				for _, cs := range el.(*ast.CaseClause).List {
// 					switch cs := cs.(type) {
// 					case *ast.StarExpr:
// 						if sexp, ok := cs.X.(*ast.SelectorExpr); ok {
// 							name = sexp.Sel.Name
// 							// name = fmt.Sprintf("%s.%s", sexp.X.(*ast.Ident).Name, sexp.X.(*ast.Ident).Name)
// 						}

// 					case *ast.SelectorExpr:
// 						name = fmt.Sprintf("%s.%s", cs.X, cs.Sel.Name)

// 					case *ast.Ident:
// 						name = cs.Name
// 					}

// 					iSl := indexSlice(name)
// 					if iSl == -1 {
// 						continue
// 					}
// 					if iSl < idx {
// 						fmt.Println("неправильный порядок switch:", fset.Position(n.Switch), name)
// 					}
// 					idx = iSl

// 				}
// 			}
// 		}

// 		return true
// 	})
// }
