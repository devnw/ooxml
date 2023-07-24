package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdEG_WrapTypeChoiceConstructor(t *testing.T) {
	v := wml.NewWdEG_WrapTypeChoice()
	if v == nil {
		t.Errorf("wml.NewWdEG_WrapTypeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdEG_WrapTypeChoice should validate: %s", err)
	}
}

func TestWdEG_WrapTypeChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdEG_WrapTypeChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdEG_WrapTypeChoice()
	xml.Unmarshal(buf, v2)
}
