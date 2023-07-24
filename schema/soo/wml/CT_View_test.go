package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ViewConstructor(t *testing.T) {
	v := wml.NewCT_View()
	if v == nil {
		t.Errorf("wml.NewCT_View must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_View should validate: %s", err)
	}
}

func TestCT_ViewMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_View()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_View()
	xml.Unmarshal(buf, v2)
}
