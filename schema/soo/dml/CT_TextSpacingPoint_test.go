package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextSpacingPointConstructor(t *testing.T) {
	v := dml.NewCT_TextSpacingPoint()
	if v == nil {
		t.Errorf("dml.NewCT_TextSpacingPoint must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextSpacingPoint should validate: %s", err)
	}
}

func TestCT_TextSpacingPointMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextSpacingPoint()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextSpacingPoint()
	xml.Unmarshal(buf, v2)
}
