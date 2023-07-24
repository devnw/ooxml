package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ColorFilterConstructor(t *testing.T) {
	v := sml.NewCT_ColorFilter()
	if v == nil {
		t.Errorf("sml.NewCT_ColorFilter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ColorFilter should validate: %s", err)
	}
}

func TestCT_ColorFilterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ColorFilter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ColorFilter()
	xml.Unmarshal(buf, v2)
}
