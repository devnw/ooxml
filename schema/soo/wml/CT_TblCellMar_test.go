package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblCellMarConstructor(t *testing.T) {
	v := wml.NewCT_TblCellMar()
	if v == nil {
		t.Errorf("wml.NewCT_TblCellMar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblCellMar should validate: %s", err)
	}
}

func TestCT_TblCellMarMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblCellMar()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblCellMar()
	xml.Unmarshal(buf, v2)
}
