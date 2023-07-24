package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AnimationDgmBuildPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_AnimationDgmBuildProperties()
	if v == nil {
		t.Errorf("dml.NewCT_AnimationDgmBuildProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AnimationDgmBuildProperties should validate: %s", err)
	}
}

func TestCT_AnimationDgmBuildPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AnimationDgmBuildProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AnimationDgmBuildProperties()
	xml.Unmarshal(buf, v2)
}
