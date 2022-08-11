package main

import (
	"time"

	xlst "github.com/ivahaev/go-xlsx-templater"
)

var ctx = map[string]interface{}{
	"name":           time.Now(),
	"nameHeader":     "Item name",
	"quantityHeader": "Quantity",
	"nm": map[string]interface{}{
		"one": "1",
		"two": "2",
	},
	"items": []map[string]interface{}{
		{
			"name":     "Pen",
			"quantity": 2,
		},
		{
			"name":     "Pencil",
			"quantity": 1,
		},
		{
			"name":     "Condom",
			"quantity": 12,
		},
		{
			"name":     "Beer",
			"quantity": 24,
		},
	},
	"rng": map[string]interface{}{
		"one": []map[string]interface{}{
			{
				"name":     "Pen",
				"quantity": 2,
			},
			{
				"name":     "Pencil",
				"quantity": 1,
			},
			{
				"name":     "Condom",
				"quantity": 12,
			},
			{
				"name":     "Beer",
				"quantity": 24,
			},
		},
	},
}

// var ctx = map[string]interface{}{
// 	"name": "TEST",
// 	"data": []map[string]interface{}{{
// 		"num":    1,
// 		"sample": "test_1",
// 		"prm": []map[string]interface{}{
// 			{
// 				"mode":  "tray",
// 				"param": "pH",
// 				"val":   "123",
// 			},
// 			{
// 				"mode":  "manual",
// 				"param": "PO2",
// 				"val":   "234",
// 			},
// 			{
// 				"mode":  "manual",
// 				"param": "Glu",
// 				"val":   234,
// 			},
// 		},
// 	}, {
// 		"num":    2,
// 		"sample": "test_2",
// 		"prm": []map[string]interface{}{
// 			{
// 				"mode":  "tray",
// 				"param": "pH",
// 				"val":   "123",
// 			},
// 			{
// 				"mode":  "manual",
// 				"param": "PO2",
// 				"val":   "234",
// 			},
// 			{
// 				"mode":  "manual",
// 				"param": "Glu",
// 				"val":   234,
// 			},
// 		},
// 	}},
// }

func main() {
	doc := xlst.New()
	err := doc.ReadTemplate("./02.xlsx")
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
