package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTemplateListConstructor(t *testing.T) {
	v := pml.NewCT_TLTemplateList()
	if v == nil {
		t.Errorf("pml.NewCT_TLTemplateList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTemplateList should validate: %s", err)
	}
}

func TestCT_TLTemplateListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTemplateList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTemplateList()
	xml.Unmarshal(buf, v2)
}
