package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_SoftEdgesEffectConstructor(t *testing.T) {
	v := dml.NewCT_SoftEdgesEffect()
	if v == nil {
		t.Errorf("dml.NewCT_SoftEdgesEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_SoftEdgesEffect should validate: %s", err)
	}
}

func TestCT_SoftEdgesEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_SoftEdgesEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_SoftEdgesEffect()
	xml.Unmarshal(buf, v2)
}
