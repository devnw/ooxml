package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalRowConstructor(t *testing.T) {
	v := sml.NewCT_ExternalRow()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalRow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalRow should validate: %s", err)
	}
}

func TestCT_ExternalRowMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalRow()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalRow()
	xml.Unmarshal(buf, v2)
}
