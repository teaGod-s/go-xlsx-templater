package main

import (
	"github.com/ivahaev/go-xlsx-templater"
)

func main() {
	doc := xlst.New()
	doc.ReadTemplate("./template.xlsx")
	ctx := map[string]interface{}{
		"name":           "Github User",
		"nameHeader":     "Item name",
		"quantityHeader": "Quantity",
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
	}
	err := doc.Render(ctx)
	if err != nil {
		panic(err)
	}
	err = doc.CopyPicture(1, `{"x_scale": 1.0, "y_scale": 1.0}`, "D1")
	if err != nil {
		panic(err)
	}
	err = doc.Save("./report.xlsx")
	if err != nil {
		panic(err)
	}
}
