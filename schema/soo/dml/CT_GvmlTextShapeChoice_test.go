package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlTextShapeChoiceConstructor(t *testing.T) {
	v := dml.NewCT_GvmlTextShapeChoice()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlTextShapeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlTextShapeChoice should validate: %s", err)
	}
}

func TestCT_GvmlTextShapeChoiceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlTextShapeChoice()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlTextShapeChoice()
	xml.Unmarshal(buf, v2)
}
