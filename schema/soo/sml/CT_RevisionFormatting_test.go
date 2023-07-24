package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionFormattingConstructor(t *testing.T) {
	v := sml.NewCT_RevisionFormatting()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionFormatting must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionFormatting should validate: %s", err)
	}
}

func TestCT_RevisionFormattingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionFormatting()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionFormatting()
	xml.Unmarshal(buf, v2)
}
