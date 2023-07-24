package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AnimationDgmElementConstructor(t *testing.T) {
	v := dml.NewCT_AnimationDgmElement()
	if v == nil {
		t.Errorf("dml.NewCT_AnimationDgmElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AnimationDgmElement should validate: %s", err)
	}
}

func TestCT_AnimationDgmElementMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AnimationDgmElement()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AnimationDgmElement()
	xml.Unmarshal(buf, v2)
}
