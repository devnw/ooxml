package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_RelativeOffsetEffectConstructor(t *testing.T) {
	v := dml.NewCT_RelativeOffsetEffect()
	if v == nil {
		t.Errorf("dml.NewCT_RelativeOffsetEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_RelativeOffsetEffect should validate: %s", err)
	}
}

func TestCT_RelativeOffsetEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_RelativeOffsetEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_RelativeOffsetEffect()
	xml.Unmarshal(buf, v2)
}
