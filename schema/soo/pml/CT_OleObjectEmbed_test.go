package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OleObjectEmbedConstructor(t *testing.T) {
	v := pml.NewCT_OleObjectEmbed()
	if v == nil {
		t.Errorf("pml.NewCT_OleObjectEmbed must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OleObjectEmbed should validate: %s", err)
	}
}

func TestCT_OleObjectEmbedMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OleObjectEmbed()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OleObjectEmbed()
	xml.Unmarshal(buf, v2)
}
