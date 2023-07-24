package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_HandlesConstructor(t *testing.T) {
	v := vml.NewCT_Handles()
	if v == nil {
		t.Errorf("vml.NewCT_Handles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Handles should validate: %s", err)
	}
}

func TestCT_HandlesMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Handles()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Handles()
	xml.Unmarshal(buf, v2)
}
