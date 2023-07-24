package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_StretchInfoPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_StretchInfoProperties()
	if v == nil {
		t.Errorf("dml.NewCT_StretchInfoProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_StretchInfoProperties should validate: %s", err)
	}
}

func TestCT_StretchInfoPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_StretchInfoProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_StretchInfoProperties()
	xml.Unmarshal(buf, v2)
}
