package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FormatConstructor(t *testing.T) {
	v := sml.NewCT_Format()
	if v == nil {
		t.Errorf("sml.NewCT_Format must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Format should validate: %s", err)
	}
}

func TestCT_FormatMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Format()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Format()
	xml.Unmarshal(buf, v2)
}
