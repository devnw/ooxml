package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_CustomGeometry2DConstructor(t *testing.T) {
	v := dml.NewCT_CustomGeometry2D()
	if v == nil {
		t.Errorf("dml.NewCT_CustomGeometry2D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_CustomGeometry2D should validate: %s", err)
	}
}

func TestCT_CustomGeometry2DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_CustomGeometry2D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_CustomGeometry2D()
	xml.Unmarshal(buf, v2)
}
