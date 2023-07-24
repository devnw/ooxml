package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_RectConstructor(t *testing.T) {
	v := vml.NewCT_Rect()
	if v == nil {
		t.Errorf("vml.NewCT_Rect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Rect should validate: %s", err)
	}
}

func TestCT_RectMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Rect()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Rect()
	xml.Unmarshal(buf, v2)
}
