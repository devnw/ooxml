package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_LineStyleListConstructor(t *testing.T) {
	v := dml.NewCT_LineStyleList()
	if v == nil {
		t.Errorf("dml.NewCT_LineStyleList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_LineStyleList should validate: %s", err)
	}
}

func TestCT_LineStyleListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_LineStyleList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_LineStyleList()
	xml.Unmarshal(buf, v2)
}
