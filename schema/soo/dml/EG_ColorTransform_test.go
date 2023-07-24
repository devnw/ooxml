package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_ColorTransformConstructor(t *testing.T) {
	v := dml.NewEG_ColorTransform()
	if v == nil {
		t.Errorf("dml.NewEG_ColorTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_ColorTransform should validate: %s", err)
	}
}

func TestEG_ColorTransformMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_ColorTransform()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_ColorTransform()
	xml.Unmarshal(buf, v2)
}
