package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MetadataBlocksConstructor(t *testing.T) {
	v := sml.NewCT_MetadataBlocks()
	if v == nil {
		t.Errorf("sml.NewCT_MetadataBlocks must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MetadataBlocks should validate: %s", err)
	}
}

func TestCT_MetadataBlocksMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MetadataBlocks()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MetadataBlocks()
	xml.Unmarshal(buf, v2)
}
