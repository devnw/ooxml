package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MarkupRangeConstructor(t *testing.T) {
	v := wml.NewCT_MarkupRange()
	if v == nil {
		t.Errorf("wml.NewCT_MarkupRange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MarkupRange should validate: %s", err)
	}
}

func TestCT_MarkupRangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MarkupRange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MarkupRange()
	xml.Unmarshal(buf, v2)
}
