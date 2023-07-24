package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_FillConstructor(t *testing.T) {
	v := vml.NewAG_Fill()
	if v == nil {
		t.Errorf("vml.NewAG_Fill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_Fill should validate: %s", err)
	}
}

func TestAG_FillMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_Fill()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_Fill()
	xml.Unmarshal(buf, v2)
}
