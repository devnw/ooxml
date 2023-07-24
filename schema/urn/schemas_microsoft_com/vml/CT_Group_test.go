package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_GroupConstructor(t *testing.T) {
	v := vml.NewCT_Group()
	if v == nil {
		t.Errorf("vml.NewCT_Group must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Group should validate: %s", err)
	}
}

func TestCT_GroupMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Group()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Group()
	xml.Unmarshal(buf, v2)
}
