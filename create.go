package main

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func pdfBuilder() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.OutputFileAndClose("./dump/dump.pdf")
}
