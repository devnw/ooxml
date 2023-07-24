package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableRowConstructor(t *testing.T) {
	v := dml.NewCT_TableRow()
	if v == nil {
		t.Errorf("dml.NewCT_TableRow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableRow should validate: %s", err)
	}
}

func TestCT_TableRowMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableRow()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableRow()
	xml.Unmarshal(buf, v2)
}
