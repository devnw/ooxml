package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DxfConstructor(t *testing.T) {
	v := sml.NewCT_Dxf()
	if v == nil {
		t.Errorf("sml.NewCT_Dxf must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Dxf should validate: %s", err)
	}
}

func TestCT_DxfMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Dxf()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Dxf()
	xml.Unmarshal(buf, v2)
}
