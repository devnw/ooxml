package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestTblStyleLstConstructor(t *testing.T) {
	v := dml.NewTblStyleLst()
	if v == nil {
		t.Errorf("dml.NewTblStyleLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.TblStyleLst should validate: %s", err)
	}
}

func TestTblStyleLstMarshalUnmarshal(t *testing.T) {
	v := dml.NewTblStyleLst()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewTblStyleLst()
	xml.Unmarshal(buf, v2)
}
