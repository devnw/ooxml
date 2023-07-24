package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimVariantBooleanValConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimVariantBooleanVal()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimVariantBooleanVal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimVariantBooleanVal should validate: %s", err)
	}
}

func TestCT_TLAnimVariantBooleanValMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimVariantBooleanVal()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimVariantBooleanVal()
	xml.Unmarshal(buf, v2)
}
