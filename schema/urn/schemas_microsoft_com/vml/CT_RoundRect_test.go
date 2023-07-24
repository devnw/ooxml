package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_RoundRectConstructor(t *testing.T) {
	v := vml.NewCT_RoundRect()
	if v == nil {
		t.Errorf("vml.NewCT_RoundRect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_RoundRect should validate: %s", err)
	}
}

func TestCT_RoundRectMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_RoundRect()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_RoundRect()
	xml.Unmarshal(buf, v2)
}
