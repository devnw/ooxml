package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DynamicFilterConstructor(t *testing.T) {
	v := sml.NewCT_DynamicFilter()
	if v == nil {
		t.Errorf("sml.NewCT_DynamicFilter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DynamicFilter should validate: %s", err)
	}
}

func TestCT_DynamicFilterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DynamicFilter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DynamicFilter()
	xml.Unmarshal(buf, v2)
}
