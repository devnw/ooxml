package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlConnectorNonVisualConstructor(t *testing.T) {
	v := dml.NewCT_GvmlConnectorNonVisual()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlConnectorNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlConnectorNonVisual should validate: %s", err)
	}
}

func TestCT_GvmlConnectorNonVisualMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlConnectorNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlConnectorNonVisual()
	xml.Unmarshal(buf, v2)
}
