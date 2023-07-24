package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TablePartStyleConstructor(t *testing.T) {
	v := dml.NewCT_TablePartStyle()
	if v == nil {
		t.Errorf("dml.NewCT_TablePartStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TablePartStyle should validate: %s", err)
	}
}

func TestCT_TablePartStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TablePartStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TablePartStyle()
	xml.Unmarshal(buf, v2)
}
