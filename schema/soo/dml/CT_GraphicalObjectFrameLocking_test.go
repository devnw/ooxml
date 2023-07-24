package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GraphicalObjectFrameLockingConstructor(t *testing.T) {
	v := dml.NewCT_GraphicalObjectFrameLocking()
	if v == nil {
		t.Errorf("dml.NewCT_GraphicalObjectFrameLocking must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GraphicalObjectFrameLocking should validate: %s", err)
	}
}

func TestCT_GraphicalObjectFrameLockingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GraphicalObjectFrameLocking()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GraphicalObjectFrameLocking()
	xml.Unmarshal(buf, v2)
}
