package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WebPrConstructor(t *testing.T) {
	v := sml.NewCT_WebPr()
	if v == nil {
		t.Errorf("sml.NewCT_WebPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WebPr should validate: %s", err)
	}
}

func TestCT_WebPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WebPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WebPr()
	xml.Unmarshal(buf, v2)
}
