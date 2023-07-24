package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_StylesConstructor(t *testing.T) {
	v := wml.NewCT_Styles()
	if v == nil {
		t.Errorf("wml.NewCT_Styles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Styles should validate: %s", err)
	}
}

func TestCT_StylesMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Styles()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Styles()
	xml.Unmarshal(buf, v2)
}
