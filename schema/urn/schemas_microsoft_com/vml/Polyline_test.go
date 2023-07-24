package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestPolylineConstructor(t *testing.T) {
	v := vml.NewPolyline()
	if v == nil {
		t.Errorf("vml.NewPolyline must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Polyline should validate: %s", err)
	}
}

func TestPolylineMarshalUnmarshal(t *testing.T) {
	v := vml.NewPolyline()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewPolyline()
	xml.Unmarshal(buf, v2)
}
