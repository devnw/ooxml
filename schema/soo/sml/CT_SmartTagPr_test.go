package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SmartTagPrConstructor(t *testing.T) {
	v := sml.NewCT_SmartTagPr()
	if v == nil {
		t.Errorf("sml.NewCT_SmartTagPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SmartTagPr should validate: %s", err)
	}
}

func TestCT_SmartTagPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SmartTagPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SmartTagPr()
	xml.Unmarshal(buf, v2)
}
