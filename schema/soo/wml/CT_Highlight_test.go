package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_HighlightConstructor(t *testing.T) {
	v := wml.NewCT_Highlight()
	if v == nil {
		t.Errorf("wml.NewCT_Highlight must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Highlight should validate: %s", err)
	}
}

func TestCT_HighlightMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Highlight()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Highlight()
	xml.Unmarshal(buf, v2)
}
