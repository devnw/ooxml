package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_RelativeRectConstructor(t *testing.T) {
	v := dml.NewCT_RelativeRect()
	if v == nil {
		t.Errorf("dml.NewCT_RelativeRect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_RelativeRect should validate: %s", err)
	}
}

func TestCT_RelativeRectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_RelativeRect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_RelativeRect()
	xml.Unmarshal(buf, v2)
}
