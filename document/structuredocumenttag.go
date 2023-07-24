package document

import "go.devnw.com/ooxml/schema/soo/wml"

// StructuredDocumentTag are a tagged bit of content in a document.
type StructuredDocumentTag struct {
	d *Document
	x *wml.CT_SdtBlock
}

// Paragraphs returns the paragraphs within a structured document tag.
func (s StructuredDocumentTag) Paragraphs() []Paragraph {
	if s.x.SdtContent == nil {
		return nil
	}
	ret := []Paragraph{}
	for _, p := range s.x.SdtContent.P {
		ret = append(ret, Paragraph{s.d, p})
	}
	return ret
}
