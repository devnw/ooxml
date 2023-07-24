package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AdjustHandleListConstructor(t *testing.T) {
	v := dml.NewCT_AdjustHandleList()
	if v == nil {
		t.Errorf("dml.NewCT_AdjustHandleList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AdjustHandleList should validate: %s", err)
	}
}

func TestCT_AdjustHandleListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AdjustHandleList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AdjustHandleList()
	xml.Unmarshal(buf, v2)
}
