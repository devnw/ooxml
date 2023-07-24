package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextShapeAutofitConstructor(t *testing.T) {
	v := dml.NewCT_TextShapeAutofit()
	if v == nil {
		t.Errorf("dml.NewCT_TextShapeAutofit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextShapeAutofit should validate: %s", err)
	}
}

func TestCT_TextShapeAutofitMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextShapeAutofit()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextShapeAutofit()
	xml.Unmarshal(buf, v2)
}
