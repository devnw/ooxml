package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MetadataTypesConstructor(t *testing.T) {
	v := sml.NewCT_MetadataTypes()
	if v == nil {
		t.Errorf("sml.NewCT_MetadataTypes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MetadataTypes should validate: %s", err)
	}
}

func TestCT_MetadataTypesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MetadataTypes()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MetadataTypes()
	xml.Unmarshal(buf, v2)
}
