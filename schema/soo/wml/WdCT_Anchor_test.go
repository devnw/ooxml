package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_AnchorConstructor(t *testing.T) {
	v := wml.NewWdCT_Anchor()
	if v == nil {
		t.Errorf("wml.NewWdCT_Anchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_Anchor should validate: %s", err)
	}
}

func TestWdCT_AnchorMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_Anchor()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_Anchor()
	xml.Unmarshal(buf, v2)
}
