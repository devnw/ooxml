package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ShapeDefaultsConstructor(t *testing.T) {
	v := wml.NewCT_ShapeDefaults()
	if v == nil {
		t.Errorf("wml.NewCT_ShapeDefaults must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ShapeDefaults should validate: %s", err)
	}
}

func TestCT_ShapeDefaultsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ShapeDefaults()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ShapeDefaults()
	xml.Unmarshal(buf, v2)
}
