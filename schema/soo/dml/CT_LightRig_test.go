package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_LightRigConstructor(t *testing.T) {
	v := dml.NewCT_LightRig()
	if v == nil {
		t.Errorf("dml.NewCT_LightRig must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_LightRig should validate: %s", err)
	}
}

func TestCT_LightRigMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_LightRig()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_LightRig()
	xml.Unmarshal(buf, v2)
}
