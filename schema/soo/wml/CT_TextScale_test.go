package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TextScaleConstructor(t *testing.T) {
	v := wml.NewCT_TextScale()
	if v == nil {
		t.Errorf("wml.NewCT_TextScale must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TextScale should validate: %s", err)
	}
}

func TestCT_TextScaleMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TextScale()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TextScale()
	xml.Unmarshal(buf, v2)
}
