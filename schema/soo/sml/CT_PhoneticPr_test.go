package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PhoneticPrConstructor(t *testing.T) {
	v := sml.NewCT_PhoneticPr()
	if v == nil {
		t.Errorf("sml.NewCT_PhoneticPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PhoneticPr should validate: %s", err)
	}
}

func TestCT_PhoneticPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PhoneticPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PhoneticPr()
	xml.Unmarshal(buf, v2)
}
