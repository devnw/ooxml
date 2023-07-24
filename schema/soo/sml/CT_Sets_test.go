package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SetsConstructor(t *testing.T) {
	v := sml.NewCT_Sets()
	if v == nil {
		t.Errorf("sml.NewCT_Sets must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Sets should validate: %s", err)
	}
}

func TestCT_SetsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Sets()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Sets()
	xml.Unmarshal(buf, v2)
}
