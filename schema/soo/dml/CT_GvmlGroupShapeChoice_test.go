package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlGroupShapeChoiceConstructor(t *testing.T) {
	v := dml.NewCT_GvmlGroupShapeChoice()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlGroupShapeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlGroupShapeChoice should validate: %s", err)
	}
}

func TestCT_GvmlGroupShapeChoiceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlGroupShapeChoice()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlGroupShapeChoice()
	xml.Unmarshal(buf, v2)
}
