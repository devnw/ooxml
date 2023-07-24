package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_FConstructor(t *testing.T) {
	v := vml.NewCT_F()
	if v == nil {
		t.Errorf("vml.NewCT_F must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_F should validate: %s", err)
	}
}

func TestCT_FMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_F()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_F()
	xml.Unmarshal(buf, v2)
}
