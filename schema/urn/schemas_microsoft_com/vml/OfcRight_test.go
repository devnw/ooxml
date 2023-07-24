package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcRightConstructor(t *testing.T) {
	v := vml.NewOfcRight()
	if v == nil {
		t.Errorf("vml.NewOfcRight must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcRight should validate: %s", err)
	}
}

func TestOfcRightMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcRight()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcRight()
	xml.Unmarshal(buf, v2)
}
