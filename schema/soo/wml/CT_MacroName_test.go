package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MacroNameConstructor(t *testing.T) {
	v := wml.NewCT_MacroName()
	if v == nil {
		t.Errorf("wml.NewCT_MacroName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MacroName should validate: %s", err)
	}
}

func TestCT_MacroNameMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MacroName()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MacroName()
	xml.Unmarshal(buf, v2)
}
