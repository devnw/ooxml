package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MdxMetadataConstructor(t *testing.T) {
	v := sml.NewCT_MdxMetadata()
	if v == nil {
		t.Errorf("sml.NewCT_MdxMetadata must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MdxMetadata should validate: %s", err)
	}
}

func TestCT_MdxMetadataMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MdxMetadata()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MdxMetadata()
	xml.Unmarshal(buf, v2)
}
