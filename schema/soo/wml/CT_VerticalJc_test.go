package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_VerticalJcConstructor(t *testing.T) {
	v := wml.NewCT_VerticalJc()
	if v == nil {
		t.Errorf("wml.NewCT_VerticalJc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_VerticalJc should validate: %s", err)
	}
}

func TestCT_VerticalJcMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_VerticalJc()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_VerticalJc()
	xml.Unmarshal(buf, v2)
}
