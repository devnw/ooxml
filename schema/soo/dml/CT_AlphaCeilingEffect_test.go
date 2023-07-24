package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AlphaCeilingEffectConstructor(t *testing.T) {
	v := dml.NewCT_AlphaCeilingEffect()
	if v == nil {
		t.Errorf("dml.NewCT_AlphaCeilingEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AlphaCeilingEffect should validate: %s", err)
	}
}

func TestCT_AlphaCeilingEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AlphaCeilingEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AlphaCeilingEffect()
	xml.Unmarshal(buf, v2)
}
