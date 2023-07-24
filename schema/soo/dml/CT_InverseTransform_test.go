package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_InverseTransformConstructor(t *testing.T) {
	v := dml.NewCT_InverseTransform()
	if v == nil {
		t.Errorf("dml.NewCT_InverseTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_InverseTransform should validate: %s", err)
	}
}

func TestCT_InverseTransformMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_InverseTransform()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_InverseTransform()
	xml.Unmarshal(buf, v2)
}
