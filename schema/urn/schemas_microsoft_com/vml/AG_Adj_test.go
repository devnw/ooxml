package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_AdjConstructor(t *testing.T) {
	v := vml.NewAG_Adj()
	if v == nil {
		t.Errorf("vml.NewAG_Adj must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_Adj should validate: %s", err)
	}
}

func TestAG_AdjMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_Adj()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_Adj()
	xml.Unmarshal(buf, v2)
}
