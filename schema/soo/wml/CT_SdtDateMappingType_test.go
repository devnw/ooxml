package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtDateMappingTypeConstructor(t *testing.T) {
	v := wml.NewCT_SdtDateMappingType()
	if v == nil {
		t.Errorf("wml.NewCT_SdtDateMappingType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtDateMappingType should validate: %s", err)
	}
}

func TestCT_SdtDateMappingTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtDateMappingType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtDateMappingType()
	xml.Unmarshal(buf, v2)
}
