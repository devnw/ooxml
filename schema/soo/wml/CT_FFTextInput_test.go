package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFTextInputConstructor(t *testing.T) {
	v := wml.NewCT_FFTextInput()
	if v == nil {
		t.Errorf("wml.NewCT_FFTextInput must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFTextInput should validate: %s", err)
	}
}

func TestCT_FFTextInputMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFTextInput()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFTextInput()
	xml.Unmarshal(buf, v2)
}
