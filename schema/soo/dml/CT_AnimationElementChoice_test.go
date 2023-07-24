package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AnimationElementChoiceConstructor(t *testing.T) {
	v := dml.NewCT_AnimationElementChoice()
	if v == nil {
		t.Errorf("dml.NewCT_AnimationElementChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AnimationElementChoice should validate: %s", err)
	}
}

func TestCT_AnimationElementChoiceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AnimationElementChoice()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AnimationElementChoice()
	xml.Unmarshal(buf, v2)
}
