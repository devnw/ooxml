package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FutureMetadataConstructor(t *testing.T) {
	v := sml.NewCT_FutureMetadata()
	if v == nil {
		t.Errorf("sml.NewCT_FutureMetadata must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FutureMetadata should validate: %s", err)
	}
}

func TestCT_FutureMetadataMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FutureMetadata()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FutureMetadata()
	xml.Unmarshal(buf, v2)
}
