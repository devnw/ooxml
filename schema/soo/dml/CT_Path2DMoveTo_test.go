package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Path2DMoveToConstructor(t *testing.T) {
	v := dml.NewCT_Path2DMoveTo()
	if v == nil {
		t.Errorf("dml.NewCT_Path2DMoveTo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Path2DMoveTo should validate: %s", err)
	}
}

func TestCT_Path2DMoveToMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Path2DMoveTo()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Path2DMoveTo()
	xml.Unmarshal(buf, v2)
}
