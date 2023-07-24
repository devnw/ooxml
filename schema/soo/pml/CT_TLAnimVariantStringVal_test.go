package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimVariantStringValConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimVariantStringVal()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimVariantStringVal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimVariantStringVal should validate: %s", err)
	}
}

func TestCT_TLAnimVariantStringValMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimVariantStringVal()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimVariantStringVal()
	xml.Unmarshal(buf, v2)
}
