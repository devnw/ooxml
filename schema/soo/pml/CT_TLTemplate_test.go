package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTemplateConstructor(t *testing.T) {
	v := pml.NewCT_TLTemplate()
	if v == nil {
		t.Errorf("pml.NewCT_TLTemplate must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTemplate should validate: %s", err)
	}
}

func TestCT_TLTemplateMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTemplate()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTemplate()
	xml.Unmarshal(buf, v2)
}
