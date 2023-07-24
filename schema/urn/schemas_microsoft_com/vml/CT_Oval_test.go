package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_OvalConstructor(t *testing.T) {
	v := vml.NewCT_Oval()
	if v == nil {
		t.Errorf("vml.NewCT_Oval must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Oval should validate: %s", err)
	}
}

func TestCT_OvalMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Oval()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Oval()
	xml.Unmarshal(buf, v2)
}
