package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Scene3DConstructor(t *testing.T) {
	v := dml.NewCT_Scene3D()
	if v == nil {
		t.Errorf("dml.NewCT_Scene3D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Scene3D should validate: %s", err)
	}
}

func TestCT_Scene3DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Scene3D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Scene3D()
	xml.Unmarshal(buf, v2)
}
