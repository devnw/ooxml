package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GraphicalObjectDataConstructor(t *testing.T) {
	v := dml.NewCT_GraphicalObjectData()
	if v == nil {
		t.Errorf("dml.NewCT_GraphicalObjectData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GraphicalObjectData should validate: %s", err)
	}
}

func TestCT_GraphicalObjectDataMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GraphicalObjectData()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GraphicalObjectData()
	xml.Unmarshal(buf, v2)
}
