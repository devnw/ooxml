package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CustomerDataConstructor(t *testing.T) {
	v := pml.NewCT_CustomerData()
	if v == nil {
		t.Errorf("pml.NewCT_CustomerData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CustomerData should validate: %s", err)
	}
}

func TestCT_CustomerDataMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CustomerData()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CustomerData()
	xml.Unmarshal(buf, v2)
}
