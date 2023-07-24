package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ShapeLockingConstructor(t *testing.T) {
	v := dml.NewCT_ShapeLocking()
	if v == nil {
		t.Errorf("dml.NewCT_ShapeLocking must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ShapeLocking should validate: %s", err)
	}
}

func TestCT_ShapeLockingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ShapeLocking()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ShapeLocking()
	xml.Unmarshal(buf, v2)
}
