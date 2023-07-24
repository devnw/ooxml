package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TextAlignmentConstructor(t *testing.T) {
	v := wml.NewCT_TextAlignment()
	if v == nil {
		t.Errorf("wml.NewCT_TextAlignment must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TextAlignment should validate: %s", err)
	}
}

func TestCT_TextAlignmentMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TextAlignment()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TextAlignment()
	xml.Unmarshal(buf, v2)
}
