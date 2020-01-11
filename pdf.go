package generate_pdf

type PDF struct {
	name     string
	pdf_type string
	page     int
	font     string
}

func PDFinit(name string, pdf_type string, page int, font string) PDF {
	pdf := PDF{
		name:     name,
		pdf_type: pdf_type,
		page:     page,
		font:     font,
	}

	return pdf
}
