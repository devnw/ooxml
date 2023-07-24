package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ZoomConstructor(t *testing.T) {
	v := wml.NewCT_Zoom()
	if v == nil {
		t.Errorf("wml.NewCT_Zoom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Zoom should validate: %s", err)
	}
}

func TestCT_ZoomMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Zoom()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Zoom()
	xml.Unmarshal(buf, v2)
}
