package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BlurEffectConstructor(t *testing.T) {
	v := dml.NewCT_BlurEffect()
	if v == nil {
		t.Errorf("dml.NewCT_BlurEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BlurEffect should validate: %s", err)
	}
}

func TestCT_BlurEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BlurEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BlurEffect()
	xml.Unmarshal(buf, v2)
}
