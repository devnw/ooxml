package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ObjectChoiceConstructor(t *testing.T) {
	v := wml.NewCT_ObjectChoice()
	if v == nil {
		t.Errorf("wml.NewCT_ObjectChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ObjectChoice should validate: %s", err)
	}
}

func TestCT_ObjectChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ObjectChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ObjectChoice()
	xml.Unmarshal(buf, v2)
}
