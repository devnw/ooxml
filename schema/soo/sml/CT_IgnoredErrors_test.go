package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_IgnoredErrorsConstructor(t *testing.T) {
	v := sml.NewCT_IgnoredErrors()
	if v == nil {
		t.Errorf("sml.NewCT_IgnoredErrors must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_IgnoredErrors should validate: %s", err)
	}
}

func TestCT_IgnoredErrorsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_IgnoredErrors()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_IgnoredErrors()
	xml.Unmarshal(buf, v2)
}
