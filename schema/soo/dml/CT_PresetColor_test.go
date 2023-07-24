package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PresetColorConstructor(t *testing.T) {
	v := dml.NewCT_PresetColor()
	if v == nil {
		t.Errorf("dml.NewCT_PresetColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PresetColor should validate: %s", err)
	}
}

func TestCT_PresetColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PresetColor()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PresetColor()
	xml.Unmarshal(buf, v2)
}
