package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_Text3DConstructor(t *testing.T) {
	v := dml.NewEG_Text3D()
	if v == nil {
		t.Errorf("dml.NewEG_Text3D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_Text3D should validate: %s", err)
	}
}

func TestEG_Text3DMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_Text3D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_Text3D()
	xml.Unmarshal(buf, v2)
}
