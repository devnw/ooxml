package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_DefaultShapeDefinitionConstructor(t *testing.T) {
	v := dml.NewCT_DefaultShapeDefinition()
	if v == nil {
		t.Errorf("dml.NewCT_DefaultShapeDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_DefaultShapeDefinition should validate: %s", err)
	}
}

func TestCT_DefaultShapeDefinitionMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_DefaultShapeDefinition()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_DefaultShapeDefinition()
	xml.Unmarshal(buf, v2)
}
