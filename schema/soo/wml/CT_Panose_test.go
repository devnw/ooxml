package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PanoseConstructor(t *testing.T) {
	v := wml.NewCT_Panose()
	if v == nil {
		t.Errorf("wml.NewCT_Panose must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Panose should validate: %s", err)
	}
}

func TestCT_PanoseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Panose()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Panose()
	xml.Unmarshal(buf, v2)
}
