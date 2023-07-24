package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalSheetDataSetConstructor(t *testing.T) {
	v := sml.NewCT_ExternalSheetDataSet()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalSheetDataSet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalSheetDataSet should validate: %s", err)
	}
}

func TestCT_ExternalSheetDataSetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalSheetDataSet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalSheetDataSet()
	xml.Unmarshal(buf, v2)
}
