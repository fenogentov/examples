package main

import (
	xlst "github.com/ivahaev/go-xlsx-templater"
)

var ctx = map[string]interface{}{
	"sample": "1234567",
	"parameter": []map[string]interface{}{
		{
			"pH": "123",
		},
		{
			"PO2": "234",
		},
	},
}

func main() {
	doc := xlst.New()
	err := doc.ReadTemplate("./00.xlsx")
	if err != nil {
		panic(err)
	}

	err = doc.Render(ctx)
	if err != nil {
		panic(err)
	}

	err = doc.Save("./report.xlsx")
	if err != nil {
		panic(err)
	}
}
