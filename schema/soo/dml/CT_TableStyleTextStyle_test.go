package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableStyleTextStyleConstructor(t *testing.T) {
	v := dml.NewCT_TableStyleTextStyle()
	if v == nil {
		t.Errorf("dml.NewCT_TableStyleTextStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableStyleTextStyle should validate: %s", err)
	}
}

func TestCT_TableStyleTextStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableStyleTextStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableStyleTextStyle()
	xml.Unmarshal(buf, v2)
}
