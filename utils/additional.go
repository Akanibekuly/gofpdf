package utils

func GetInvoice() Invoice {
	return Invoice{
		ID:              "7",
		IDdate:          "1 февраля",
		Date:            "17 декабрь 2020 года",
		Postocshik:      "(полностью прописью)",
		IINiAddress:     "ИИН ..., Республика Казахстан",
		IIK:             "KZ... в АО '...', БИК ....",
		Dogovor:         "Без договора",
		UslovyaOplati:   "безналичный расчет",
		PunktNazn:       "",
		SposobOtpravki:  "99 (Прочие)",
		Nakladnaya:      "",
		GruzOtpravitel:  "ИИН ...",
		GrusPolychatel:  "БИН: ...",
		Poluchatel:      "(полностью прописью)",
		BINiAddress:     "БИН: 1..., Республика Казахстан, г. Нур-Султан,",
		IIKpolychatelya: "KZ..., в банке АО '...', БИК ....",
	}
}

func GetCompanyInfo() Company {
	return Company{
		Name: "Ivozprovider",
		Logo: "avatar.jpg",
	}
}
