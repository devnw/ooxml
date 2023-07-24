package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_AutoCaptionConstructor(t *testing.T) {
	v := wml.NewCT_AutoCaption()
	if v == nil {
		t.Errorf("wml.NewCT_AutoCaption must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_AutoCaption should validate: %s", err)
	}
}

func TestCT_AutoCaptionMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_AutoCaption()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_AutoCaption()
	xml.Unmarshal(buf, v2)
}
