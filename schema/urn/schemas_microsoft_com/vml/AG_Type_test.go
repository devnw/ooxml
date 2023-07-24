package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_TypeConstructor(t *testing.T) {
	v := vml.NewAG_Type()
	if v == nil {
		t.Errorf("vml.NewAG_Type must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_Type should validate: %s", err)
	}
}

func TestAG_TypeMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_Type()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_Type()
	xml.Unmarshal(buf, v2)
}
