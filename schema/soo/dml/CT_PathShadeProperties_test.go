package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PathShadePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_PathShadeProperties()
	if v == nil {
		t.Errorf("dml.NewCT_PathShadeProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PathShadeProperties should validate: %s", err)
	}
}

func TestCT_PathShadePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PathShadeProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PathShadeProperties()
	xml.Unmarshal(buf, v2)
}
