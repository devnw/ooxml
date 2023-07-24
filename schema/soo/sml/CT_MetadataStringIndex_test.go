package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MetadataStringIndexConstructor(t *testing.T) {
	v := sml.NewCT_MetadataStringIndex()
	if v == nil {
		t.Errorf("sml.NewCT_MetadataStringIndex must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MetadataStringIndex should validate: %s", err)
	}
}

func TestCT_MetadataStringIndexMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MetadataStringIndex()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MetadataStringIndex()
	xml.Unmarshal(buf, v2)
}
