package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ConnectionConstructor(t *testing.T) {
	v := dml.NewCT_Connection()
	if v == nil {
		t.Errorf("dml.NewCT_Connection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Connection should validate: %s", err)
	}
}

func TestCT_ConnectionMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Connection()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Connection()
	xml.Unmarshal(buf, v2)
}
