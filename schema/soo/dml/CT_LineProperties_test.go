package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_LinePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_LineProperties()
	if v == nil {
		t.Errorf("dml.NewCT_LineProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_LineProperties should validate: %s", err)
	}
}

func TestCT_LinePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_LineProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_LineProperties()
	xml.Unmarshal(buf, v2)
}
