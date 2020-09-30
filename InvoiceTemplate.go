package main

import (
	"log"

	"./example"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddFont("Helvetica", "", "./assets/helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	tr := pdf.UnicodeTranslatorFromDescriptor("./assets/cp1251")

	// Creating header
	pdf.SetTopMargin(30) //set margin
	pdf.SetHeaderFuncMode(func() {
		pdf.Image(example.ImageFile("logo.png"), 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 15)
		pdf.Cell(80, 0, "")
		pdf.CellFormat(30, 10, "Title", "1", 0, "C", false, 0, "")
		pdf.Ln(20)
	}, true)

	pdf.Cell(15, 50, tr("Акжол и Була ываыаыва"))
	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		log.Println(err)
	}
}
