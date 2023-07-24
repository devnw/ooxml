package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MetadataBlockConstructor(t *testing.T) {
	v := sml.NewCT_MetadataBlock()
	if v == nil {
		t.Errorf("sml.NewCT_MetadataBlock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MetadataBlock should validate: %s", err)
	}
}

func TestCT_MetadataBlockMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MetadataBlock()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MetadataBlock()
	xml.Unmarshal(buf, v2)
}
