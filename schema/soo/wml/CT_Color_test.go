package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ColorConstructor(t *testing.T) {
	v := wml.NewCT_Color()
	if v == nil {
		t.Errorf("wml.NewCT_Color must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Color should validate: %s", err)
	}
}

func TestCT_ColorMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Color()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Color()
	xml.Unmarshal(buf, v2)
}
