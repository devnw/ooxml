package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ConditionalFormatsConstructor(t *testing.T) {
	v := sml.NewCT_ConditionalFormats()
	if v == nil {
		t.Errorf("sml.NewCT_ConditionalFormats must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ConditionalFormats should validate: %s", err)
	}
}

func TestCT_ConditionalFormatsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ConditionalFormats()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ConditionalFormats()
	xml.Unmarshal(buf, v2)
}
