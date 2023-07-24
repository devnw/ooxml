package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TableStyleListConstructor(t *testing.T) {
	v := dml.NewCT_TableStyleList()
	if v == nil {
		t.Errorf("dml.NewCT_TableStyleList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableStyleList should validate: %s", err)
	}
}

func TestCT_TableStyleListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableStyleList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableStyleList()
	xml.Unmarshal(buf, v2)
}
