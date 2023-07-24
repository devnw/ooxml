package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_EntryConstructor(t *testing.T) {
	v := vml.NewOfcCT_Entry()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Entry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Entry should validate: %s", err)
	}
}

func TestOfcCT_EntryMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Entry()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Entry()
	xml.Unmarshal(buf, v2)
}
