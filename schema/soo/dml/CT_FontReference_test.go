package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_FontReferenceConstructor(t *testing.T) {
	v := dml.NewCT_FontReference()
	if v == nil {
		t.Errorf("dml.NewCT_FontReference must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FontReference should validate: %s", err)
	}
}

func TestCT_FontReferenceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FontReference()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FontReference()
	xml.Unmarshal(buf, v2)
}
