package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Scale2DConstructor(t *testing.T) {
	v := dml.NewCT_Scale2D()
	if v == nil {
		t.Errorf("dml.NewCT_Scale2D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Scale2D should validate: %s", err)
	}
}

func TestCT_Scale2DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Scale2D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Scale2D()
	xml.Unmarshal(buf, v2)
}
