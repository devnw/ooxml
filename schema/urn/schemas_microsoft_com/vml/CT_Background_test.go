package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_BackgroundConstructor(t *testing.T) {
	v := vml.NewCT_Background()
	if v == nil {
		t.Errorf("vml.NewCT_Background must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Background should validate: %s", err)
	}
}

func TestCT_BackgroundMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Background()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Background()
	xml.Unmarshal(buf, v2)
}
