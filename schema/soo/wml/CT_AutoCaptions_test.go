package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_AutoCaptionsConstructor(t *testing.T) {
	v := wml.NewCT_AutoCaptions()
	if v == nil {
		t.Errorf("wml.NewCT_AutoCaptions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_AutoCaptions should validate: %s", err)
	}
}

func TestCT_AutoCaptionsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_AutoCaptions()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_AutoCaptions()
	xml.Unmarshal(buf, v2)
}
