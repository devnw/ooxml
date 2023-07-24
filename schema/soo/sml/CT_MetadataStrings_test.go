package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MetadataStringsConstructor(t *testing.T) {
	v := sml.NewCT_MetadataStrings()
	if v == nil {
		t.Errorf("sml.NewCT_MetadataStrings must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MetadataStrings should validate: %s", err)
	}
}

func TestCT_MetadataStringsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MetadataStrings()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MetadataStrings()
	xml.Unmarshal(buf, v2)
}
