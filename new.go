package main

import (
	"fmt"

	"github.com/Akanibekuly/gofpdf/example"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	err := GeneratePdf("Template.pdf")
	if err != nil {
		panic(err)
	}
}

func GeneratePdf(filename string) error {
	const offset = 75.0
	var companyName string
	companyName = "Название компании"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddFont("Helvetica", "", "./assets/helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 8)
	tr := pdf.UnicodeTranslatorFromDescriptor("./assets/cp1251")

	// add company name
	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(350, 5, tr(companyName), "0", 0, "CM", false, 0, "")

	// add logo image
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

	//function that draws lines
	var draw = func(x0, y0, x1, y1 float64) {
		pdf.SetDrawColor(0x33, 0x33, 0x33)
		pdf.SetLineWidth(0.4)
		pdf.MoveTo(x0, y0)
		pdf.LineTo(x1, y1)
		pdf.DrawPath("D")
	}
	// we draw main line after header
	draw(10, 30, 200, 30)

	// write title
	fontSize := 16.0
	ht := pdf.PointConvert(fontSize)
	pdf.SetFontSize(fontSize)
	pdf.Ln(ht)
	titleStr := "Счет-фактура № 1 от 7 февраля 2020 г."
	wd := pdf.GetStringWidth(titleStr) + 6
	pdf.SetX((210 - wd) / 2)
	pdf.CellFormat(wd, ht, tr(titleStr), "", 0, "C", false, 0, "")
	pdf.Ln(ht)

	//draw table information
	fontSize = 8
	pdf.SetFontSize(fontSize)
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

	//draw Invoice table
	xPos := 46.0
	for _, v := range arr {
		pdf.Ln(ht * 0.05)
		pdf.CellFormat(190, ht, tr(v), "", 0, "L", false, 0, "")
		draw(10.0, xPos, 190.0, xPos)
		xPos += 4.5
	}

	pdf.Ln(ht * 2)

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

	//рисуем конечную таблицу
	marginCell := 1. // margin of top/bottom of cell
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
	// заготовки таблицы
	cols := []float64{8.0, 25.0, 15.0, 20.0, 15.0, 30.0, 20.0, 20.0, 20.0}
	rows := [][]string{}
	rows = append(rows, header)
	for _, row := range rows {
		curx, y := pdf.GetXY()
		x := curx

		height := 14.0
		_, lineHt := pdf.GetFontSize()

		// add a new page if the height of the row doesn't fit on the page
		if pdf.GetY()+height > pageh-mbottom {
			pdf.AddPage()
			y = pdf.GetY()
		}
		fmt.Println("height", height)
		for i, txt := range row {
			width := cols[i]
			pdf.Rect(x, y, width, height, "")
			fmt.Println(txt)
			pdf.MultiCell(width, lineHt+marginCell, tr(txt), "", "C", false)
			x += width
			pdf.SetXY(x, y)
		}
		pdf.SetXY(curx, y+height)
	}
	err := pdf.OutputFileAndClose(filename)
	example.Summary(err, filename)
	return err
}

func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}
