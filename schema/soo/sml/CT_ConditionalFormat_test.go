package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ConditionalFormatConstructor(t *testing.T) {
	v := sml.NewCT_ConditionalFormat()
	if v == nil {
		t.Errorf("sml.NewCT_ConditionalFormat must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ConditionalFormat should validate: %s", err)
	}
}

func TestCT_ConditionalFormatMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ConditionalFormat()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ConditionalFormat()
	xml.Unmarshal(buf, v2)
}
