package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableStyleCellStyleConstructor(t *testing.T) {
	v := dml.NewCT_TableStyleCellStyle()
	if v == nil {
		t.Errorf("dml.NewCT_TableStyleCellStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableStyleCellStyle should validate: %s", err)
	}
}

func TestCT_TableStyleCellStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableStyleCellStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableStyleCellStyle()
	xml.Unmarshal(buf, v2)
}
