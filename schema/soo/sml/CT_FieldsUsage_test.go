package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FieldsUsageConstructor(t *testing.T) {
	v := sml.NewCT_FieldsUsage()
	if v == nil {
		t.Errorf("sml.NewCT_FieldsUsage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FieldsUsage should validate: %s", err)
	}
}

func TestCT_FieldsUsageMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FieldsUsage()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FieldsUsage()
	xml.Unmarshal(buf, v2)
}
