package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLBuildDiagramConstructor(t *testing.T) {
	v := pml.NewCT_TLBuildDiagram()
	if v == nil {
		t.Errorf("pml.NewCT_TLBuildDiagram must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLBuildDiagram should validate: %s", err)
	}
}

func TestCT_TLBuildDiagramMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLBuildDiagram()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLBuildDiagram()
	xml.Unmarshal(buf, v2)
}
