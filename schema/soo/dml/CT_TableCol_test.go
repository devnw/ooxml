package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableColConstructor(t *testing.T) {
	v := dml.NewCT_TableCol()
	if v == nil {
		t.Errorf("dml.NewCT_TableCol must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableCol should validate: %s", err)
	}
}

func TestCT_TableColMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableCol()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableCol()
	xml.Unmarshal(buf, v2)
}
