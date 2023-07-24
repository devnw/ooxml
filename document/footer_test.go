package document_test

import (
	"testing"

	"go.devnw.com/ooxml/document"
)

func TestFooterRemoveParagraph(t *testing.T) {
	doc := document.New()
	for i := 0; i < 5; i++ {
		hdr := doc.AddFooter()
		for j := 0; j < i; j++ {
			hdr.AddParagraph()
		}
	}

	if len(doc.Footers()) != 5 {
		t.Errorf("expected 5 paragraphs, found %d", len(doc.Footers()))
	}
	for i, hdr := range doc.Footers() {
		if len(hdr.Paragraphs()) != i {
			t.Errorf("expected %d paragraphs in document, found %d", i, len(hdr.Paragraphs()))
		}
		for _, p := range hdr.Paragraphs() {
			hdr.RemoveParagraph(p)
		}
		if len(hdr.Paragraphs()) != 0 {
			t.Errorf("expected no paragraphs, found %d", len(hdr.Paragraphs()))
		}
	}

}
