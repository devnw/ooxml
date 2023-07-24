package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableStyleConstructor(t *testing.T) {
	v := dml.NewCT_TableStyle()
	if v == nil {
		t.Errorf("dml.NewCT_TableStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableStyle should validate: %s", err)
	}
}

func TestCT_TableStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableStyle()
	xml.Unmarshal(buf, v2)
}
