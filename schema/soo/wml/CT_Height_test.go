package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_HeightConstructor(t *testing.T) {
	v := wml.NewCT_Height()
	if v == nil {
		t.Errorf("wml.NewCT_Height must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Height should validate: %s", err)
	}
}

func TestCT_HeightMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Height()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Height()
	xml.Unmarshal(buf, v2)
}
