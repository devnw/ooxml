package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFTextTypeConstructor(t *testing.T) {
	v := wml.NewCT_FFTextType()
	if v == nil {
		t.Errorf("wml.NewCT_FFTextType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFTextType should validate: %s", err)
	}
}

func TestCT_FFTextTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFTextType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFTextType()
	xml.Unmarshal(buf, v2)
}
