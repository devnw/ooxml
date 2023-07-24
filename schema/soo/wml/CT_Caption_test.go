package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CaptionConstructor(t *testing.T) {
	v := wml.NewCT_Caption()
	if v == nil {
		t.Errorf("wml.NewCT_Caption must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Caption should validate: %s", err)
	}
}

func TestCT_CaptionMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Caption()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Caption()
	xml.Unmarshal(buf, v2)
}
