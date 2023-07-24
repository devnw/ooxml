package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimVariantConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimVariant()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimVariant must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimVariant should validate: %s", err)
	}
}

func TestCT_TLAnimVariantMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimVariant()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimVariant()
	xml.Unmarshal(buf, v2)
}
