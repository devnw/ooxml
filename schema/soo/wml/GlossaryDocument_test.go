package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestGlossaryDocumentConstructor(t *testing.T) {
	v := wml.NewGlossaryDocument()
	if v == nil {
		t.Errorf("wml.NewGlossaryDocument must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.GlossaryDocument should validate: %s", err)
	}
}

func TestGlossaryDocumentMarshalUnmarshal(t *testing.T) {
	v := wml.NewGlossaryDocument()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewGlossaryDocument()
	xml.Unmarshal(buf, v2)
}
