package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableCellPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_TableCellProperties()
	if v == nil {
		t.Errorf("dml.NewCT_TableCellProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableCellProperties should validate: %s", err)
	}
}

func TestCT_TableCellPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableCellProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableCellProperties()
	xml.Unmarshal(buf, v2)
}
