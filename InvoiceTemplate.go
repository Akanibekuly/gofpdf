package main

import (
	"strings"

	"github.com/Akanibekuly/gofpdf/example/pkg"
	"github.com/Akanibekuly/gofpdf/example"
	
	
	"github.com/jung-kurt/gofpdf"
)

func main() {
	// err := GeneratePdf("InvoiceTemplate.pdf")
	// if err != nil {
	// 	panic(err)
	// }
	ExampleFpdf_SplitLines_tables()
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
	// countryList := make([]countryType, 0, 8)
	header := []string{
		`№
		п/п`,
		`Наименование товаров 
		(работ, услуг)`,
		`Ед. изм.`,
		`Кол-во
		(объем)`,
		`Цена
		(KZT)`,
		`Стоимость товаров 
		(работ, услуг) 
		без НДС`,
		`НДС`,
		`Всего
		стоимость
		реализации`,
		"Акциз",
	}

	width := []float64{20, 35, 20, 20, 25, 30, 35, 25, 35}
	// loadData := func(fileStr string) {
	// 	fl, err := os.Open(fileStr)
	// 	if err == nil {
	// 		scanner := bufio.NewScanner(fl)
	// 		var c countryType
	// 		for scanner.Scan() {
	// 			// Austria;Vienna;83859;8075
	// 			lineStr := scanner.Text()
	// 			list := strings.Split(lineStr, ";")
	// 			if len(list) == 4 {
	// 				c.nameStr = list[0]
	// 				c.capitalStr = list[1]
	// 				c.areaStr = list[2]
	// 				c.popStr = list[3]
	// 				countryList = append(countryList, c)
	// 			} else {
	// 				err = fmt.Errorf("error tokenizing %s", lineStr)
	// 			}
	// 		}
	// 		fl.Close()
	// 		if len(countryList) == 0 {
	// 			err = fmt.Errorf("error loading data from %s", fileStr)
	// 		}
	// 	}
	// 	if err != nil {
	// 		pdf.SetError(err)
	// 	}
	// }

	// Simple table
	basicTable := func() {
		left := (210.0 - 4*40) / 2
		pdf.SetX(left)
		for i, str := range header {
			pdf.CellFormat(width[i], 7, tr(str), "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
		// for _, c := range countryList {
		// 	pdf.SetX(left)
		// 	pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
		// 	pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
		// 	pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
		// 	pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
		// 	pdf.Ln(-1)
		// }
	}
	// loadData("./countries.txt")
	pdf.Ln(ht * 3)
	basicTable()

	return pdf.OutputFileAndClose(filename)
}

// func ExampleFpdf_CellFormat_tables() {
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	type countryType struct {
// 		nameStr, capitalStr, areaStr, popStr string
// 	}
// 	countryList := make([]countryType, 0, 8)
// 	header := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}
// 	loadData := func(fileStr string) {
// 		fl, err := os.Open(fileStr)
// 		if err == nil {
// 			scanner := bufio.NewScanner(fl)
// 			var c countryType
// 			for scanner.Scan() {
// 				// Austria;Vienna;83859;8075
// 				lineStr := scanner.Text()
// 				list := strings.Split(lineStr, ";")
// 				if len(list) == 4 {
// 					c.nameStr = list[0]
// 					c.capitalStr = list[1]
// 					c.areaStr = list[2]
// 					c.popStr = list[3]
// 					countryList = append(countryList, c)
// 				} else {
// 					err = fmt.Errorf("error tokenizing %s", lineStr)
// 				}
// 			}
// 			fl.Close()
// 			if len(countryList) == 0 {
// 				err = fmt.Errorf("error loading data from %s", fileStr)
// 			}
// 		}
// 		if err != nil {
// 			pdf.SetError(err)
// 		}
// 	}
// 	// Simple table
// 	basicTable := func() {
// 		left := (210.0 - 4*40) / 2
// 		pdf.SetX(left)
// 		for _, str := range header {
// 			pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
// 		}
// 		pdf.Ln(-1)
// 		for _, c := range countryList {
// 			pdf.SetX(left)
// 			pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
// 			pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
// 			pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
// 			pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
// 			pdf.Ln(-1)
// 		}
// 	}
// 	// Better table
// 	improvedTable := func() {
// 		// Column widths
// 		w := []float64{40.0, 35.0, 40.0, 45.0}
// 		wSum := 0.0
// 		for _, v := range w {
// 			wSum += v
// 		}
// 		left := (210 - wSum) / 2
// 		// 	Header
// 		pdf.SetX(left)
// 		for j, str := range header {
// 			pdf.CellFormat(w[j], 7, str, "1", 0, "C", false, 0, "")
// 		}
// 		pdf.Ln(-1)
// 		// Data
// 		for _, c := range countryList {
// 			pdf.SetX(left)
// 			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", false, 0, "")
// 			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", false, 0, "")
// 			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
// 				"LR", 0, "R", false, 0, "")
// 			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
// 				"LR", 0, "R", false, 0, "")
// 			pdf.Ln(-1)
// 		}
// 		pdf.SetX(left)
// 		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
// 	}
// 	// Colored table
// 	fancyTable := func() {
// 		// Colors, line width and bold font
// 		pdf.SetFillColor(255, 0, 0)
// 		pdf.SetTextColor(255, 255, 255)
// 		pdf.SetDrawColor(128, 0, 0)
// 		pdf.SetLineWidth(.3)
// 		pdf.SetFont("", "B", 0)
// 		// 	Header
// 		w := []float64{40, 35, 40, 45}
// 		wSum := 0.0
// 		for _, v := range w {
// 			wSum += v
// 		}
// 		left := (210 - wSum) / 2
// 		pdf.SetX(left)
// 		for j, str := range header {
// 			pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
// 		}
// 		pdf.Ln(-1)
// 		// Color and font restoration
// 		pdf.SetFillColor(224, 235, 255)
// 		pdf.SetTextColor(0, 0, 0)
// 		pdf.SetFont("", "", 0)
// 		// 	Data
// 		fill := false
// 		for _, c := range countryList {
// 			pdf.SetX(left)
// 			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", fill, 0, "")
// 			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", fill, 0, "")
// 			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
// 				"LR", 0, "R", fill, 0, "")
// 			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
// 				"LR", 0, "R", fill, 0, "")
// 			pdf.Ln(-1)
// 			fill = !fill
// 		}
// 		pdf.SetX(left)
// 		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
// 	}
// 	loadData("./countries.txt")
// 	pdf.SetFont("Arial", "", 14)
// 	pdf.AddPage()
// 	basicTable()
// 	pdf.AddPage()
// 	improvedTable()
// 	pdf.AddPage()
// 	fancyTable()
// 	fileStr := "Fpdf_CellFormat_tables.pdf"
// 	err := pdf.OutputFileAndClose(fileStr)
// 	example.Summary(err, fileStr)
// 	// Output:
// 	// Successfully generated pdf/Fpdf_CellFormat_tables.pdf
// }

func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}

func ExampleFpdf_SplitLines_tables() {
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
	strList := pkg.LoremList()
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
