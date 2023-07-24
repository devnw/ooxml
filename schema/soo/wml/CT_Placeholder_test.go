package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PlaceholderConstructor(t *testing.T) {
	v := wml.NewCT_Placeholder()
	if v == nil {
		t.Errorf("wml.NewCT_Placeholder must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Placeholder should validate: %s", err)
	}
}

func TestCT_PlaceholderMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Placeholder()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Placeholder()
	xml.Unmarshal(buf, v2)
}
