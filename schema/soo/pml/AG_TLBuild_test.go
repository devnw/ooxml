package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestAG_TLBuildConstructor(t *testing.T) {
	v := pml.NewAG_TLBuild()
	if v == nil {
		t.Errorf("pml.NewAG_TLBuild must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.AG_TLBuild should validate: %s", err)
	}
}

func TestAG_TLBuildMarshalUnmarshal(t *testing.T) {
	v := pml.NewAG_TLBuild()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewAG_TLBuild()
	xml.Unmarshal(buf, v2)
}
