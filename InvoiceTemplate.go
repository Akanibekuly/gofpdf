package main

import (
	// "strings"
	"github.com/Akanibekuly/gofpdf/example"
	// "github.com/Akanibekuly/gofpdf/utils"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	err := GeneratePdf("InvoiceTemplate.pdf")
	if err != nil {
		panic(err)
	}
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
	// we draw main line after header
	draw("butt", "bevel", 10, 30, 200, 30)
	fontSize := 16.0
	// write title
	pdf.SetFontSize(fontSize)
	ht := pdf.PointConvert(fontSize)
	write := func(str, align string) {
		pdf.CellFormat(190, ht, tr(str), "", 1, align, false, 0, "")
	}
	pdf.Ln(ht)
	write(tr("Счет-фактура № 1 от 7 февраля 2020 г."), "C")
	pdf.Ln(ht)
	// end title

	//  start invoisce table
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

	//draw Invoice table
	xPos := 50.0
	for i, v := range arr {
		if i != 0 {
			pdf.Ln(ht * 0.3)
		}
		write(tr(v), "L")
		draw("butt", "bevel", 10, xPos, 190, xPos)
		xPos += 4
	}

	err:=pdf.OutputFileAndClose(filename)
	example.Summary(err, filename)
	return  err


}
