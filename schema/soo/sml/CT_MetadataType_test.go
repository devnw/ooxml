package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MetadataTypeConstructor(t *testing.T) {
	v := sml.NewCT_MetadataType()
	if v == nil {
		t.Errorf("sml.NewCT_MetadataType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MetadataType should validate: %s", err)
	}
}

func TestCT_MetadataTypeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MetadataType()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MetadataType()
	xml.Unmarshal(buf, v2)
}
