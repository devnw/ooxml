package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionAutoFormattingConstructor(t *testing.T) {
	v := sml.NewCT_RevisionAutoFormatting()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionAutoFormatting must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionAutoFormatting should validate: %s", err)
	}
}

func TestCT_RevisionAutoFormattingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionAutoFormatting()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionAutoFormatting()
	xml.Unmarshal(buf, v2)
}
