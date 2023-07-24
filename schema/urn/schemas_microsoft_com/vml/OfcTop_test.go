package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcTopConstructor(t *testing.T) {
	v := vml.NewOfcTop()
	if v == nil {
		t.Errorf("vml.NewOfcTop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcTop should validate: %s", err)
	}
}

func TestOfcTopMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcTop()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcTop()
	xml.Unmarshal(buf, v2)
}
