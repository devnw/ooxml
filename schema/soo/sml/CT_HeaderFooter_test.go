package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_HeaderFooterConstructor(t *testing.T) {
	v := sml.NewCT_HeaderFooter()
	if v == nil {
		t.Errorf("sml.NewCT_HeaderFooter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_HeaderFooter should validate: %s", err)
	}
}

func TestCT_HeaderFooterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_HeaderFooter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_HeaderFooter()
	xml.Unmarshal(buf, v2)
}
