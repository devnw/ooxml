package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AnimationGraphicalObjectBuildPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_AnimationGraphicalObjectBuildProperties()
	if v == nil {
		t.Errorf("dml.NewCT_AnimationGraphicalObjectBuildProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AnimationGraphicalObjectBuildProperties should validate: %s", err)
	}
}

func TestCT_AnimationGraphicalObjectBuildPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AnimationGraphicalObjectBuildProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AnimationGraphicalObjectBuildProperties()
	xml.Unmarshal(buf, v2)
}
