package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_OleSizeConstructor(t *testing.T) {
	v := sml.NewCT_OleSize()
	if v == nil {
		t.Errorf("sml.NewCT_OleSize must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_OleSize should validate: %s", err)
	}
}

func TestCT_OleSizeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_OleSize()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_OleSize()
	xml.Unmarshal(buf, v2)
}
