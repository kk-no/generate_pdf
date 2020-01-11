package generate_pdf

type PDF struct {
	name string
	page int
	font string
}

func PDFinit(name string, page int, font string) PDF {
	pdf := PDF{
		name: name,
		page: page,
		font: font,
	}

	return pdf
}
