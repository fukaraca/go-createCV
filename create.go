package main

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"os"
)

//pdfGenerator creates the CV.pdf from buffered data which is templated HTML with provided information
func pdfGenerator(r *bytes.Buffer) error {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println("pdf couldn't be generated", err)
		return err
	}
	//prepare a new page from a buffered stream of bytes
	page := wkhtmltopdf.NewPageReader(r)
	page.EnableLocalFileAccess.Set(true)
	page.Zoom.Set(1.5)
	//append the page to PdfGenerator
	pdfg.AddPage(page)

	pdfg.SetStderr(os.Stdout)
	pdfg.Dpi.Set(300)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	err = pdfg.Create()
	if err != nil {
		log.Println("pdf couldn't be created:", err)
		return err
	}

	err = pdfg.WriteFile("./web/dump/CV.pdf")
	if err != nil {
		log.Println("pdf couldn't be writed ", err)
		return err
	}

	return nil
}

//templater simply creates HTML version of PDF to be generated and returns the data as buffered bytes
func (info *Info) templater() *bytes.Buffer {
	t := template.New("template.html")

	t, err := t.ParseFiles("./template.html")
	if err != nil {
		log.Println("template.html couldn't be parsed", err)
		return nil
	}

	body := bytes.Buffer{}
	body.Grow(1000)
	err = t.Execute(&body, gin.H{
		"fullname":   info.Fullname,
		"title":      info.Title,
		"profile":    info.Profile,
		"skills":     info.Skills,
		"references": info.References,
		"contact":    info.Contact,
		"career":     info.Career,
		"education":  info.Education,
		"pdfname":    info.Fullname + "'s CV",
		"photograph": info.photoPath,
		"addpath":    AddPath(),
	})
	if err != nil {
		log.Println("template couldn't be executed: ", err)
	}

	return &body
}

//AddPath function returns working directory in order to make wkhtmltopdf able-to-find/access to template-static-files
func AddPath() template.URL {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("current directory couldn't be get", err)
	}

	return template.URL(fmt.Sprintf("file://%s", dir))
}
