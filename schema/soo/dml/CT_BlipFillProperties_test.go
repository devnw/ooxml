package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BlipFillPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_BlipFillProperties()
	if v == nil {
		t.Errorf("dml.NewCT_BlipFillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BlipFillProperties should validate: %s", err)
	}
}

func TestCT_BlipFillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BlipFillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BlipFillProperties()
	xml.Unmarshal(buf, v2)
}
