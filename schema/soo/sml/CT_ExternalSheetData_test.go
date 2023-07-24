package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalSheetDataConstructor(t *testing.T) {
	v := sml.NewCT_ExternalSheetData()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalSheetData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalSheetData should validate: %s", err)
	}
}

func TestCT_ExternalSheetDataMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalSheetData()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalSheetData()
	xml.Unmarshal(buf, v2)
}
