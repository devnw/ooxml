package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcSkewConstructor(t *testing.T) {
	v := vml.NewOfcSkew()
	if v == nil {
		t.Errorf("vml.NewOfcSkew must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcSkew should validate: %s", err)
	}
}

func TestOfcSkewMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcSkew()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcSkew()
	xml.Unmarshal(buf, v2)
}
