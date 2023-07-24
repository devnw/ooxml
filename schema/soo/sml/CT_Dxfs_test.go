package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DxfsConstructor(t *testing.T) {
	v := sml.NewCT_Dxfs()
	if v == nil {
		t.Errorf("sml.NewCT_Dxfs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Dxfs should validate: %s", err)
	}
}

func TestCT_DxfsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Dxfs()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Dxfs()
	xml.Unmarshal(buf, v2)
}
