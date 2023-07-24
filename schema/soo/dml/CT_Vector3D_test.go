package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Vector3DConstructor(t *testing.T) {
	v := dml.NewCT_Vector3D()
	if v == nil {
		t.Errorf("dml.NewCT_Vector3D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Vector3D should validate: %s", err)
	}
}

func TestCT_Vector3DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Vector3D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Vector3D()
	xml.Unmarshal(buf, v2)
}
