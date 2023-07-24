package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_LinearShadePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_LinearShadeProperties()
	if v == nil {
		t.Errorf("dml.NewCT_LinearShadeProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_LinearShadeProperties should validate: %s", err)
	}
}

func TestCT_LinearShadePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_LinearShadeProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_LinearShadeProperties()
	xml.Unmarshal(buf, v2)
}
