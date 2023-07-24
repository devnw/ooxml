package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimVariantIntegerValConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimVariantIntegerVal()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimVariantIntegerVal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimVariantIntegerVal should validate: %s", err)
	}
}

func TestCT_TLAnimVariantIntegerValMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimVariantIntegerVal()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimVariantIntegerVal()
	xml.Unmarshal(buf, v2)
}
