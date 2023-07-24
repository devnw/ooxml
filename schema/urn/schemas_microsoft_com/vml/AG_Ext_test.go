package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_ExtConstructor(t *testing.T) {
	v := vml.NewAG_Ext()
	if v == nil {
		t.Errorf("vml.NewAG_Ext must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_Ext should validate: %s", err)
	}
}

func TestAG_ExtMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_Ext()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_Ext()
	xml.Unmarshal(buf, v2)
}
