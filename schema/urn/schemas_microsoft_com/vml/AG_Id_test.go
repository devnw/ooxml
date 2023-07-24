package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_IdConstructor(t *testing.T) {
	v := vml.NewAG_Id()
	if v == nil {
		t.Errorf("vml.NewAG_Id must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_Id should validate: %s", err)
	}
}

func TestAG_IdMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_Id()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_Id()
	xml.Unmarshal(buf, v2)
}
