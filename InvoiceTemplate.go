package main

import (
	"strings"

	"github.com/Akanibekuly/gofpdf/example"
	"github.com/Akanibekuly/gofpdf/utils"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	// err := GeneratePdf("InvoiceTemplate.pdf")
	// if err != nil {
	// 	panic(err)
	// }
	drawTable()
}

// GeneratePdf generates our pdf by adding text and images to the page
// then saving it to a file (name specified in params).
func GeneratePdf(filename string) error {
	const offset = 75.0
	var companyName string
	companyName = "Название компании"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddFont("Helvetica", "", "./assets/helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 8)
	tr := pdf.UnicodeTranslatorFromDescriptor("./assets/cp1251")

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(350, 5, tr(companyName), "0", 0, "CM", false, 0, "")

	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"avatar.jpg",
		175, 15,
		20, 12,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)
	var draw = func(cap, join string, x0, y0, x1, y1 float64) {
		// transform begin & end needed to isolate caps and joins
		pdf.SetLineCapStyle(cap)
		pdf.SetLineJoinStyle(join)
		// Draw thick line
		pdf.SetDrawColor(0x33, 0x33, 0x33)
		pdf.SetLineWidth(0.4)
		pdf.MoveTo(x0, y0)
		pdf.LineTo(x1, y1)
		pdf.DrawPath("D")

	}
	// caps := []string{"butt", "square", "round"}
	// joins := []string{"bevel", "miter", "round"}
	draw("butt", "bevel", 10, 30, 200, 30)
	fontSize := 16.0
	pdf.SetFont("Helvetica", "", fontSize)
	ht := pdf.PointConvert(fontSize)
	write := func(str, align string) {
		pdf.CellFormat(190, ht, str, "", 1, align, false, 0, "")
	}
	pdf.Ln(ht)
	write(tr("Счет-фактура № 1 от 7 февраля 2020 г."), "C")
	pdf.Ln(ht)

	//  начало счет фактуры
	fontSize = 8
	pdf.SetFontSize(fontSize)

	ht = pdf.PointConvert(fontSize)
	arr := []string{
		"Дата совершения оборота:",
		"Поставщик: (полностью прописью)",
		"ИИН и адрес места нахождения поставщика: ИИН …, Республика Казахстан",
		"ИИК поставщика: KZ... в АО '...', БИК ….",
		"Договор (контракт) на поставку товаров (работ, услуг): Без договора",
		"Условия оплаты по договору (контракту): безналичный расчет",
		"Пункт назначения поставляемых товаров (работ, услуг): ",
		"Поставка товаров (работ,услуг) осуществлена по доверенности: Без доверенности",
		"Способ отправления: 99 (Прочие)",
		"Товарно-транспортная накладная: ",
		"Грузоотправитель:   ИИН …",
		"Грузополучатель: БИН: …",
		"Получатель: (полностью прописью)",
		"БИН и адрес места нахождения получателя: БИН: 1…, Республика Казахстан, г. Нур-Султан, ",
		"ИИК получателя: KZ..., в банке АО '...', БИК ….",
	}

	xPos := 50.0
	for i, v := range arr {
		if i != 0 {
			pdf.Ln(ht * 0.3)
		}
		write(tr(v), "L")
		draw("butt", "bevel", 10, xPos, 190, xPos)
		xPos += 4
	}

	type countryType struct {
		nameStr, capitalStr, areaStr, popStr string
	}

	drawTable()
	return pdf.OutputFileAndClose(filename)
}

func drawTable() {
	const (
		colCount = 3
		colWd    = 60.0
		marginH  = 15.0
		lineHt   = 5.5
		cellGap  = 2.0
	)
	// var colStrList [colCount]string
	type cellType struct {
		str  string
		list [][]byte
		ht   float64
	}
	var (
		cellList [colCount]cellType
		cell     cellType
	)
	pdf := gofpdf.New("P", "mm", "A4", "") // 210 x 297
	header := [colCount]string{"Column A", "Column B", "Column C"}
	alignList := [colCount]string{"L", "C", "R"}
	strList := utils.LoremList()
	pdf.SetMargins(marginH, 15, marginH)
	pdf.SetFont("Arial", "", 14)
	pdf.AddPage()

	// Headers
	pdf.SetTextColor(224, 224, 224)
	pdf.SetFillColor(64, 64, 64)
	for colJ := 0; colJ < colCount; colJ++ {
		pdf.CellFormat(colWd, 10, header[colJ], "1", 0, "CM", true, 0, "")
	}
	pdf.Ln(-1)
	pdf.SetTextColor(24, 24, 24)
	pdf.SetFillColor(255, 255, 255)

	// Rows
	y := pdf.GetY()
	count := 0
	for rowJ := 0; rowJ < 2; rowJ++ {
		maxHt := lineHt
		// Cell height calculation loop
		for colJ := 0; colJ < colCount; colJ++ {
			count++
			if count > len(strList) {
				count = 1
			}
			cell.str = strings.Join(strList[0:count], " ")
			cell.list = pdf.SplitLines([]byte(cell.str), colWd-cellGap-cellGap)
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[colJ] = cell
		}
		// Cell render loop
		x := marginH
		for colJ := 0; colJ < colCount; colJ++ {
			pdf.Rect(x, y, colWd, maxHt+cellGap+cellGap, "D")
			cell = cellList[colJ]
			cellY := y + cellGap + (maxHt-cell.ht)/2
			for splitJ := 0; splitJ < len(cell.list); splitJ++ {
				pdf.SetXY(x+cellGap, cellY)
				pdf.CellFormat(colWd-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
					alignList[colJ], false, 0, "")
				cellY += lineHt
			}
			x += colWd
		}
		y += maxHt + cellGap + cellGap
	}

	fileStr := "Fpdf_SplitLines_tables.pdf"
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/Fpdf_SplitLines_tables.pdf
}

// func loremList() []string {
// 	return []string{
// 		"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod " +
// 			"tempor incididunt ut labore et dolore magna aliqua.",
// 		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
// 			"aliquip ex ea commodo consequat.",
// 		"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum " +
// 			"dolore eu fugiat nulla pariatur.",
// 		"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui " +
// 			"officia deserunt mollit anim id est laborum.",
// 	}
// }
