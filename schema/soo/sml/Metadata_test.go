package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestMetadataConstructor(t *testing.T) {
	v := sml.NewMetadata()
	if v == nil {
		t.Errorf("sml.NewMetadata must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Metadata should validate: %s", err)
	}
}

func TestMetadataMarshalUnmarshal(t *testing.T) {
	v := sml.NewMetadata()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewMetadata()
	xml.Unmarshal(buf, v2)
}
