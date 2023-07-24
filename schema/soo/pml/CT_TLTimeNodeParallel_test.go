package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTimeNodeParallelConstructor(t *testing.T) {
	v := pml.NewCT_TLTimeNodeParallel()
	if v == nil {
		t.Errorf("pml.NewCT_TLTimeNodeParallel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTimeNodeParallel should validate: %s", err)
	}
}

func TestCT_TLTimeNodeParallelMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTimeNodeParallel()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTimeNodeParallel()
	xml.Unmarshal(buf, v2)
}
