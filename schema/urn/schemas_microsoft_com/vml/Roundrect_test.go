package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestRoundrectConstructor(t *testing.T) {
	v := vml.NewRoundrect()
	if v == nil {
		t.Errorf("vml.NewRoundrect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Roundrect should validate: %s", err)
	}
}

func TestRoundrectMarshalUnmarshal(t *testing.T) {
	v := vml.NewRoundrect()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewRoundrect()
	xml.Unmarshal(buf, v2)
}
