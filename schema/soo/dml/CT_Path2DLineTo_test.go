package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Path2DLineToConstructor(t *testing.T) {
	v := dml.NewCT_Path2DLineTo()
	if v == nil {
		t.Errorf("dml.NewCT_Path2DLineTo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Path2DLineTo should validate: %s", err)
	}
}

func TestCT_Path2DLineToMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Path2DLineTo()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Path2DLineTo()
	xml.Unmarshal(buf, v2)
}
