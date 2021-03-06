package main

import (
	"fmt"
	"strconv"

	"github.com/Akanibekuly/gofpdf/utils"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	err := GeneratePdf("Template.pdf")
	if err != nil {
		panic(err)
	}
}

func GeneratePdf(filename string) error {
	compInfo := utils.GetCompanyInfo()
	Invoice := utils.GetInvoice()
	const offset = 75.0
	var companyName string
	companyName = compInfo.Name
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
		compInfo.Logo,
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
	draw(15, 30, 200, 30)

	// write title
	fontSize := 16.0
	ht := pdf.PointConvert(fontSize)
	pdf.SetFontSize(fontSize)
	pdf.Ln(ht)
	titleStr := "Счет-фактура № " + Invoice.ID + " от " + Invoice.IDdate
	wd := pdf.GetStringWidth(titleStr) + 6
	pdf.SetX((210 - wd) / 2)
	pdf.CellFormat(wd, ht, tr(titleStr), "", 0, "C", false, 0, "")
	pdf.Ln(ht)

	//draw table information
	fontSize = 9
	pdf.SetFontSize(fontSize)

	arr := []string{
		"Дата совершения оборота: " + Invoice.Date,
		"Поставщик: " + Invoice.Postocshik,
		"ИИН и адрес места нахождения поставщика: " + Invoice.IINiAddress,
		"ИИК поставщика: " + Invoice.IIK,
		"Договор (контракт) на поставку товаров (работ, услуг): " + Invoice.Dogovor,
		"Условия оплаты по договору (контракту): " + Invoice.UslovyaOplati,
		"Пункт назначения поставляемых товаров (работ, услуг): " + Invoice.PunktNazn,
		"Поставка товаров (работ,услуг): " + Invoice.Postavka,
		"Способ отправления: " + Invoice.SposobOtpravki,
		"Товарно-транспортная накладная: " + Invoice.Nakladnaya,
		"Грузоотправитель: " + Invoice.GruzOtpravitel,
		"Грузополучатель: " + Invoice.GrusPolychatel,
		"Получатель: " + Invoice.Poluchatel,
		"БИН и адрес места нахождения получателя: " + Invoice.BINiAddress,
		"ИИК получателя: " + Invoice.IIKpolychatelya,
	}

	//draw Invoice table

	xPos := 46.0
	for _, v := range arr {
		pdf.Ln(ht * 0.05)
		pdf.SetX(15)
		pdf.CellFormat(190, ht, tr(v), "", 0, "L", false, 0, "")
		draw(15.0, xPos, 190.0, xPos)
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

	curx, y := pdf.GetXY()
	curx += 5
	pdf.SetXY(curx, y)
	x := curx
	fmt.Println("x=", curx)
	var drawTables = func(cols []float64, rows [][]string, height float64) {
		for _, row := range rows {
			y = pdf.GetY()
			x = curx
			_, lineHt := pdf.GetFontSize()
			// add a new page if the height of the row doesn't fit on the page
			if pdf.GetY()+height > pageh-mbottom {
				pdf.AddPage()
				y = pdf.GetY()
			}

			for i, txt := range row {
				width := cols[i]
				pdf.Rect(x, y, width, height, "")
				pdf.MultiCell(width, lineHt+marginCell, tr(txt), "", "CM", false)
				if txt == "НДС" || txt == "Акциз" {
					draw(x, y+height/2, x+width, y+height/2)
					draw(x+width/2, y+height/2, x+width/2, y+height)
					pdf.SetXY(x, y+height/2)
					pdf.MultiCell(width/2, lineHt+marginCell, tr("Ставка"), "", "CM", false)
					pdf.SetXY(x, y+height)
					pdf.SetXY(x+width/2, y+height/2)
					pdf.MultiCell(width/2, lineHt+marginCell, tr("Сумма"), "", "CM", false)
					pdf.SetXY(x+width/2, y+height)
				}
				x += width
				pdf.SetXY(x, y)
			}
			y += height
			pdf.SetXY(x, y)
		}

	}

	// заготовки таблицы
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
	cols := []float64{8.0, 25.0, 15.0, 20.0, 15.0, 30.0, 25.0, 20.0, 25.0}
	rows := [][]string{}
	rows = append(rows, header)
	drawTables(cols, rows, 14)

	var temp = []string{}
	for i := 0; i < 11; i++ {
		temp = append(temp, strconv.Itoa(i+1))
	}
	pdf.SetXY(curx, y)
	rows = [][]string{}
	cols = []float64{8.0, 25.0, 15.0, 20.0, 15.0, 30.0, 12.5, 12.5, 20.0, 12.5, 12.5}
	rows = append(rows, temp)
	drawTables(cols, rows, 5)

	height := 14.0
	y = pdf.GetY()
	pdf.SetFontSize(10)
	pdf.SetXY(curx, y+height*2)
	pdf.Cell(50, 10, tr("Руководитель:"))
	draw(curx, y+height*2+12, curx+76, y+height*2+12)
	pdf.SetXY(curx+100, y+height*2)
	pdf.Cell(50, 10, tr("ВЫДАЛ (ответственное лицо поставщика)"))
	draw(curx+100, y+height*2+12, curx+185, y+height*2+12)

	pdf.SetFontSize(8)
	pdf.SetXY(curx, y+height*2+12)
	pdf.CellFormat(75, 4, tr("(Ф.И.О подпись)"), "", 1, "C", false, 0, "")
	pdf.SetXY(curx+100, y+height*2+12)
	pdf.CellFormat(90, 4, tr("(должность)"), "", 1, "C", false, 0, "")

	pdf.SetFontSize(13)
	pdf.SetXY(curx+83, y+height*2+14)
	pdf.CellFormat(10, 8, tr("МП"), "1", 1, "C", false, 0, "")

	pdf.SetFontSize(10)
	y = pdf.GetY()
	y -= 5
	pdf.SetXY(curx, y)
	pdf.Cell(50, 10, tr("Главный бухгалтер: Не предусмотрен"))
	draw(curx, y+12, curx+76, y+12)
	draw(curx+100, y+12, curx+185, y+12)
	pdf.SetFontSize(8)
	pdf.SetXY(curx, y+12)
	pdf.CellFormat(75, 4, tr("(Ф.И.О подпись)"), "", 1, "C", false, 0, "")
	pdf.SetXY(curx+100, y+12)
	pdf.CellFormat(90, 4, tr("(Ф.И.О подпись)"), "", 1, "C", false, 0, "")
	err := pdf.OutputFileAndClose(filename)
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
