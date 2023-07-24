package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MetadataConstructor(t *testing.T) {
	v := sml.NewCT_Metadata()
	if v == nil {
		t.Errorf("sml.NewCT_Metadata must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Metadata should validate: %s", err)
	}
}

func TestCT_MetadataMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Metadata()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Metadata()
	xml.Unmarshal(buf, v2)
}
