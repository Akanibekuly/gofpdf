package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Akanibekuly/gofpdf/example"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	ExampleFpdf_MultiCell()
	// pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", "A4", "")
	// pdf.AddPage()
	// pdf.AddFont("Helvetica", "", "../assets/helvetica_1251.json")
	// pdf.SetFont("Helvetica", "", 8)
	// tr := pdf.UnicodeTranslatorFromDescriptor("../assets/cp1251")
	// pdf.Cell(40, 10, tr("Hello World!  Акжол"))
	// fileStr := "basic.pdf"
	// err := pdf.OutputFileAndClose(fileStr)
	// example.Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/basic.pdf
}

func ExampleFpdf_MultiCell() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	titleStr := "20000 Leagues Under the Seas"
	pdf.SetTitle(titleStr, false)
	pdf.SetAuthor("Jules Verne", false)
	pdf.SetHeaderFunc(func() {
		// Arial bold 15
		pdf.SetFont("Arial", "B", 15)
		// Calculate width of title and position
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetX((210 - wd) / 2)
		// Colors of frame, background and text
		pdf.SetDrawColor(0, 80, 180)
		pdf.SetFillColor(230, 230, 0)
		pdf.SetTextColor(220, 50, 50)
		// Thickness of frame (1 mm)
		pdf.SetLineWidth(1)
		// Title
		pdf.CellFormat(wd, 9, titleStr, "1", 1, "C", true, 0, "")
		// Line break
		pdf.Ln(10)
	})
	pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		pdf.SetY(-15)
		// Arial italic 8
		pdf.SetFont("Arial", "I", 8)
		// Text color in gray
		pdf.SetTextColor(128, 128, 128)
		// Page number
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	chapterTitle := func(chapNum int, titleStr string) {
		// 	// Arial 12
		pdf.SetFont("Arial", "", 12)
		// Background color
		pdf.SetFillColor(200, 220, 255)
		// Title
		pdf.CellFormat(0, 6, fmt.Sprintf("Chapter %d : %s", chapNum, titleStr),
			"", 1, "L", true, 0, "")
		// Line break
		pdf.Ln(4)
	}
	chapterBody := func(fileStr string) {
		// Read text file
		txtStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}
		// Times 12
		pdf.SetFont("Times", "", 12)
		// Output justified text
		pdf.MultiCell(0, 5, string(txtStr), "", "", false)
		// Line break
		pdf.Ln(-1)
		// Mention in italics
		pdf.SetFont("", "I", 0)
		pdf.Cell(0, 5, "(end of excerpt)")
	}
	printChapter := func(chapNum int, titleStr, fileStr string) {
		pdf.AddPage()
		chapterTitle(chapNum, titleStr)
		chapterBody(fileStr)
	}
	printChapter(1, "A RUNAWAY REEF", "20k_c1.txt")
	printChapter(2, "THE PROS AND CONS", "20k_c2.txt")
	fileStr := "Fpdf_MultiCell.pdf"
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/Fpdf_MultiCell.pdf
}
