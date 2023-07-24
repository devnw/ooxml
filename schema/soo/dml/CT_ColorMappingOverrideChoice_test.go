package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ColorMappingOverrideChoiceConstructor(t *testing.T) {
	v := dml.NewCT_ColorMappingOverrideChoice()
	if v == nil {
		t.Errorf("dml.NewCT_ColorMappingOverrideChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorMappingOverrideChoice should validate: %s", err)
	}
}

func TestCT_ColorMappingOverrideChoiceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorMappingOverrideChoice()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorMappingOverrideChoice()
	xml.Unmarshal(buf, v2)
}
