package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextListStyleConstructor(t *testing.T) {
	v := dml.NewCT_TextListStyle()
	if v == nil {
		t.Errorf("dml.NewCT_TextListStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextListStyle should validate: %s", err)
	}
}

func TestCT_TextListStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextListStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextListStyle()
	xml.Unmarshal(buf, v2)
}
