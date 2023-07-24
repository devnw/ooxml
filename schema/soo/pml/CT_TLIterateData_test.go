package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLIterateDataConstructor(t *testing.T) {
	v := pml.NewCT_TLIterateData()
	if v == nil {
		t.Errorf("pml.NewCT_TLIterateData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLIterateData should validate: %s", err)
	}
}

func TestCT_TLIterateDataMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLIterateData()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLIterateData()
	xml.Unmarshal(buf, v2)
}
