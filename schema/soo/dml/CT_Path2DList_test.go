package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Path2DListConstructor(t *testing.T) {
	v := dml.NewCT_Path2DList()
	if v == nil {
		t.Errorf("dml.NewCT_Path2DList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Path2DList should validate: %s", err)
	}
}

func TestCT_Path2DListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Path2DList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Path2DList()
	xml.Unmarshal(buf, v2)
}
