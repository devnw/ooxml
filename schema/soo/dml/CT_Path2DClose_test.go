package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Path2DCloseConstructor(t *testing.T) {
	v := dml.NewCT_Path2DClose()
	if v == nil {
		t.Errorf("dml.NewCT_Path2DClose must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Path2DClose should validate: %s", err)
	}
}

func TestCT_Path2DCloseMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Path2DClose()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Path2DClose()
	xml.Unmarshal(buf, v2)
}
