package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FormatsConstructor(t *testing.T) {
	v := sml.NewCT_Formats()
	if v == nil {
		t.Errorf("sml.NewCT_Formats must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Formats should validate: %s", err)
	}
}

func TestCT_FormatsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Formats()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Formats()
	xml.Unmarshal(buf, v2)
}
