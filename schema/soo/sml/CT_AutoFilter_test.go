package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_AutoFilterConstructor(t *testing.T) {
	v := sml.NewCT_AutoFilter()
	if v == nil {
		t.Errorf("sml.NewCT_AutoFilter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_AutoFilter should validate: %s", err)
	}
}

func TestCT_AutoFilterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_AutoFilter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_AutoFilter()
	xml.Unmarshal(buf, v2)
}
