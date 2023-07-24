package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestTblConstructor(t *testing.T) {
	v := dml.NewTbl()
	if v == nil {
		t.Errorf("dml.NewTbl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.Tbl should validate: %s", err)
	}
}

func TestTblMarshalUnmarshal(t *testing.T) {
	v := dml.NewTbl()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewTbl()
	xml.Unmarshal(buf, v2)
}
