package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DimensionsConstructor(t *testing.T) {
	v := sml.NewCT_Dimensions()
	if v == nil {
		t.Errorf("sml.NewCT_Dimensions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Dimensions should validate: %s", err)
	}
}

func TestCT_DimensionsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Dimensions()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Dimensions()
	xml.Unmarshal(buf, v2)
}
