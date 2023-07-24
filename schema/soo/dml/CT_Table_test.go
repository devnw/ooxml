package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableConstructor(t *testing.T) {
	v := dml.NewCT_Table()
	if v == nil {
		t.Errorf("dml.NewCT_Table must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Table should validate: %s", err)
	}
}

func TestCT_TableMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Table()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Table()
	xml.Unmarshal(buf, v2)
}
