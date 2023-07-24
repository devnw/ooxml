package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_StyleConstructor(t *testing.T) {
	v := vml.NewAG_Style()
	if v == nil {
		t.Errorf("vml.NewAG_Style must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_Style should validate: %s", err)
	}
}

func TestAG_StyleMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_Style()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_Style()
	xml.Unmarshal(buf, v2)
}
