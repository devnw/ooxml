package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_GlossaryDocumentConstructor(t *testing.T) {
	v := wml.NewCT_GlossaryDocument()
	if v == nil {
		t.Errorf("wml.NewCT_GlossaryDocument must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_GlossaryDocument should validate: %s", err)
	}
}

func TestCT_GlossaryDocumentMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_GlossaryDocument()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_GlossaryDocument()
	xml.Unmarshal(buf, v2)
}
