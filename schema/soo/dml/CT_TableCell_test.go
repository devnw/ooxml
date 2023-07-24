package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableCellConstructor(t *testing.T) {
	v := dml.NewCT_TableCell()
	if v == nil {
		t.Errorf("dml.NewCT_TableCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableCell should validate: %s", err)
	}
}

func TestCT_TableCellMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableCell()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableCell()
	xml.Unmarshal(buf, v2)
}
