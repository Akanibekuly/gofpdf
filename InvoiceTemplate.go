package main

import (
	"fmt"

	"github.com/Akanibekuly/gofpdf/example"
	"github.com/jung-kurt/gofpdf"
)

// "github.com/Akanibekuly/gofpdf/"

// func main() {
// 	ExampleFpdf_Rect()
// }

func ExampleFpdf_Rect() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddFont("Helvetica", "", "./assets/helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 8)
	tr := pdf.UnicodeTranslatorFromDescriptor("./assets/cp1251")

	marginCell := 2. // margin of top/bottom of cell
	pagew, pageh := pdf.GetPageSize()
	mleft, mright, _, mbottom := pdf.GetMargins()
	fmt.Println(pagew, pageh, mleft, mright)
	header := []string{
		"№ п/п",
		"Наименование товаров (работ, услуг)",
		"Ед. изм.",
		"Кол-во (объем)",
		"Цена (KZT)",
		"Стоимость товаров (работ, услуг) без НДС",
		"НДС",
		"Всего стоимость реализации",
		"Акциз",
	}
	cols := []float64{8.0, 25.0, 15.0, 20.0, 20.0, 30.0, 30.0, 20.0, 30.0}
	rows := [][]string{}

	rows = append(rows, header)

	for _, row := range rows {
		curx, y := pdf.GetXY()
		x := curx

		height := 0.
		_, lineHt := pdf.GetFontSize()

		for i, txt := range row {
			lines := pdf.SplitLines([]byte(txt), cols[i])
			h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
			if h > height {
				height = h
			}
		}
		// add a new page if the height of the row doesn't fit on the page
		if pdf.GetY()+height > pageh-mbottom {
			pdf.AddPage()
			y = pdf.GetY()
		}
		for i, txt := range row {
			width := cols[i]
			pdf.Rect(x, y, width, height, "")
			pdf.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
			x += width
			pdf.SetXY(x, y)
		}
		pdf.SetXY(curx, y+height)
	}
	fileStr := "Fpdf_WrappedTableCells.pdf"
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/Fpdf_WrappedTableCells.pdf
}
