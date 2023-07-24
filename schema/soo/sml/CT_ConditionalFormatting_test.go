package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ConditionalFormattingConstructor(t *testing.T) {
	v := sml.NewCT_ConditionalFormatting()
	if v == nil {
		t.Errorf("sml.NewCT_ConditionalFormatting must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ConditionalFormatting should validate: %s", err)
	}
}

func TestCT_ConditionalFormattingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ConditionalFormatting()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ConditionalFormatting()
	xml.Unmarshal(buf, v2)
}
