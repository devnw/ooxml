package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLGraphicalObjectBuildConstructor(t *testing.T) {
	v := pml.NewCT_TLGraphicalObjectBuild()
	if v == nil {
		t.Errorf("pml.NewCT_TLGraphicalObjectBuild must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLGraphicalObjectBuild should validate: %s", err)
	}
}

func TestCT_TLGraphicalObjectBuildMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLGraphicalObjectBuild()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLGraphicalObjectBuild()
	xml.Unmarshal(buf, v2)
}
