package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SectTypeConstructor(t *testing.T) {
	v := wml.NewCT_SectType()
	if v == nil {
		t.Errorf("wml.NewCT_SectType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SectType should validate: %s", err)
	}
}

func TestCT_SectTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SectType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SectType()
	xml.Unmarshal(buf, v2)
}
