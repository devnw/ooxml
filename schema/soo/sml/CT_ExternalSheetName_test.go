package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalSheetNameConstructor(t *testing.T) {
	v := sml.NewCT_ExternalSheetName()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalSheetName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalSheetName should validate: %s", err)
	}
}

func TestCT_ExternalSheetNameMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalSheetName()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalSheetName()
	xml.Unmarshal(buf, v2)
}
