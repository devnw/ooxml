package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimVariantFloatValConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimVariantFloatVal()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimVariantFloatVal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimVariantFloatVal should validate: %s", err)
	}
}

func TestCT_TLAnimVariantFloatValMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimVariantFloatVal()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimVariantFloatVal()
	xml.Unmarshal(buf, v2)
}
