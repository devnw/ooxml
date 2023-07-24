package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AnimationChartBuildPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_AnimationChartBuildProperties()
	if v == nil {
		t.Errorf("dml.NewCT_AnimationChartBuildProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AnimationChartBuildProperties should validate: %s", err)
	}
}

func TestCT_AnimationChartBuildPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AnimationChartBuildProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AnimationChartBuildProperties()
	xml.Unmarshal(buf, v2)
}
