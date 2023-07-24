package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_HeaderFooterConstructor(t *testing.T) {
	v := pml.NewCT_HeaderFooter()
	if v == nil {
		t.Errorf("pml.NewCT_HeaderFooter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_HeaderFooter should validate: %s", err)
	}
}

func TestCT_HeaderFooterMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_HeaderFooter()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_HeaderFooter()
	xml.Unmarshal(buf, v2)
}
