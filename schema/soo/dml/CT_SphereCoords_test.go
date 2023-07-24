package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_SphereCoordsConstructor(t *testing.T) {
	v := dml.NewCT_SphereCoords()
	if v == nil {
		t.Errorf("dml.NewCT_SphereCoords must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_SphereCoords should validate: %s", err)
	}
}

func TestCT_SphereCoordsMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_SphereCoords()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_SphereCoords()
	xml.Unmarshal(buf, v2)
}
