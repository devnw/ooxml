package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BooleanConstructor(t *testing.T) {
	v := dml.NewCT_Boolean()
	if v == nil {
		t.Errorf("dml.NewCT_Boolean must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Boolean should validate: %s", err)
	}
}

func TestCT_BooleanMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Boolean()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Boolean()
	xml.Unmarshal(buf, v2)
}
