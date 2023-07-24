package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFCheckBoxChoiceConstructor(t *testing.T) {
	v := wml.NewCT_FFCheckBoxChoice()
	if v == nil {
		t.Errorf("wml.NewCT_FFCheckBoxChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFCheckBoxChoice should validate: %s", err)
	}
}

func TestCT_FFCheckBoxChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFCheckBoxChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFCheckBoxChoice()
	xml.Unmarshal(buf, v2)
}
