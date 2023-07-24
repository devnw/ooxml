package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_SystemColorConstructor(t *testing.T) {
	v := dml.NewCT_SystemColor()
	if v == nil {
		t.Errorf("dml.NewCT_SystemColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_SystemColor should validate: %s", err)
	}
}

func TestCT_SystemColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_SystemColor()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_SystemColor()
	xml.Unmarshal(buf, v2)
}
