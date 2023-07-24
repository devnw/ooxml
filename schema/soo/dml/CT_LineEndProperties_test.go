package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_LineEndPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_LineEndProperties()
	if v == nil {
		t.Errorf("dml.NewCT_LineEndProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_LineEndProperties should validate: %s", err)
	}
}

func TestCT_LineEndPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_LineEndProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_LineEndProperties()
	xml.Unmarshal(buf, v2)
}
