package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_StylePaneFilterConstructor(t *testing.T) {
	v := wml.NewCT_StylePaneFilter()
	if v == nil {
		t.Errorf("wml.NewCT_StylePaneFilter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_StylePaneFilter should validate: %s", err)
	}
}

func TestCT_StylePaneFilterMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_StylePaneFilter()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_StylePaneFilter()
	xml.Unmarshal(buf, v2)
}
