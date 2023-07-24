package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FutureMetadataBlockConstructor(t *testing.T) {
	v := sml.NewCT_FutureMetadataBlock()
	if v == nil {
		t.Errorf("sml.NewCT_FutureMetadataBlock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FutureMetadataBlock should validate: %s", err)
	}
}

func TestCT_FutureMetadataBlockMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FutureMetadataBlock()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FutureMetadataBlock()
	xml.Unmarshal(buf, v2)
}
