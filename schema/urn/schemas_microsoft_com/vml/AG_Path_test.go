package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_PathConstructor(t *testing.T) {
	v := vml.NewAG_Path()
	if v == nil {
		t.Errorf("vml.NewAG_Path must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_Path should validate: %s", err)
	}
}

func TestAG_PathMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_Path()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_Path()
	xml.Unmarshal(buf, v2)
}
