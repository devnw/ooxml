package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FieldUsageConstructor(t *testing.T) {
	v := sml.NewCT_FieldUsage()
	if v == nil {
		t.Errorf("sml.NewCT_FieldUsage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FieldUsage should validate: %s", err)
	}
}

func TestCT_FieldUsageMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FieldUsage()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FieldUsage()
	xml.Unmarshal(buf, v2)
}
