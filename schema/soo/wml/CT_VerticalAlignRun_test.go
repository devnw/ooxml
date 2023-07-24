package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_VerticalAlignRunConstructor(t *testing.T) {
	v := wml.NewCT_VerticalAlignRun()
	if v == nil {
		t.Errorf("wml.NewCT_VerticalAlignRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_VerticalAlignRun should validate: %s", err)
	}
}

func TestCT_VerticalAlignRunMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_VerticalAlignRun()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_VerticalAlignRun()
	xml.Unmarshal(buf, v2)
}
