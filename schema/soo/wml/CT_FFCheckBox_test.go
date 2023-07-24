package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFCheckBoxConstructor(t *testing.T) {
	v := wml.NewCT_FFCheckBox()
	if v == nil {
		t.Errorf("wml.NewCT_FFCheckBox must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFCheckBox should validate: %s", err)
	}
}

func TestCT_FFCheckBoxMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFCheckBox()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFCheckBox()
	xml.Unmarshal(buf, v2)
}
