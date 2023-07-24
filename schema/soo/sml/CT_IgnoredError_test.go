package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_IgnoredErrorConstructor(t *testing.T) {
	v := sml.NewCT_IgnoredError()
	if v == nil {
		t.Errorf("sml.NewCT_IgnoredError must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_IgnoredError should validate: %s", err)
	}
}

func TestCT_IgnoredErrorMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_IgnoredError()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_IgnoredError()
	xml.Unmarshal(buf, v2)
}
