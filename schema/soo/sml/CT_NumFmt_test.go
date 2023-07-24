package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_NumFmtConstructor(t *testing.T) {
	v := sml.NewCT_NumFmt()
	if v == nil {
		t.Errorf("sml.NewCT_NumFmt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_NumFmt should validate: %s", err)
	}
}

func TestCT_NumFmtMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_NumFmt()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_NumFmt()
	xml.Unmarshal(buf, v2)
}
