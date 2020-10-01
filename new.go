package main

import (
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
		// transform begin & end needed to isolate caps and joins
		// pdf.SetLineCapStyle(cap)
		// pdf.SetLineJoinStyle(join)
		// Draw thick line
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

	err := pdf.OutputFileAndClose(filename)
	example.Summary(err, filename)
	return err
}
