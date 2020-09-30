package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddFont("Helvetica", "", "./assets/helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	tr := pdf.UnicodeTranslatorFromDescriptor("./assets/cp1251")
	pdf.Cell(15, 50, tr("Акжол и Була ываыаыва"))
	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		log.Println(err)
	}
}
