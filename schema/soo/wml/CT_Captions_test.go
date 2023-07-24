package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CaptionsConstructor(t *testing.T) {
	v := wml.NewCT_Captions()
	if v == nil {
		t.Errorf("wml.NewCT_Captions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Captions should validate: %s", err)
	}
}

func TestCT_CaptionsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Captions()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Captions()
	xml.Unmarshal(buf, v2)
}
