package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GrayscaleTransformConstructor(t *testing.T) {
	v := dml.NewCT_GrayscaleTransform()
	if v == nil {
		t.Errorf("dml.NewCT_GrayscaleTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GrayscaleTransform should validate: %s", err)
	}
}

func TestCT_GrayscaleTransformMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GrayscaleTransform()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GrayscaleTransform()
	xml.Unmarshal(buf, v2)
}
