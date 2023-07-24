package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextParagraphPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_TextParagraphProperties()
	if v == nil {
		t.Errorf("dml.NewCT_TextParagraphProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextParagraphProperties should validate: %s", err)
	}
}

func TestCT_TextParagraphPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextParagraphProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextParagraphProperties()
	xml.Unmarshal(buf, v2)
}
