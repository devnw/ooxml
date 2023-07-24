package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_FillConstructor(t *testing.T) {
	v := vml.NewCT_Fill()
	if v == nil {
		t.Errorf("vml.NewCT_Fill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Fill should validate: %s", err)
	}
}

func TestCT_FillMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Fill()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Fill()
	xml.Unmarshal(buf, v2)
}
