package main

import (
	"go.devnw.com/ooxml/document"
)

func main() {
	doc, err := document.Open("document.docx")
	if err != nil {
		panic(err)
	}
	doc.SetStrict(false) // document will be saved as Word document (this is a default option for new files)
	doc.SaveToFile("conformance_transitional.docx")
	doc.SetStrict(true) // document will be saved in the Strict mode
	doc.SaveToFile("conformance_strict.docx")
}
