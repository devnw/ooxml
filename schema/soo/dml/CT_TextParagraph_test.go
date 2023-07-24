package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextParagraphConstructor(t *testing.T) {
	v := dml.NewCT_TextParagraph()
	if v == nil {
		t.Errorf("dml.NewCT_TextParagraph must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextParagraph should validate: %s", err)
	}
}

func TestCT_TextParagraphMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextParagraph()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextParagraph()
	xml.Unmarshal(buf, v2)
}
