package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AlphaBiLevelEffectConstructor(t *testing.T) {
	v := dml.NewCT_AlphaBiLevelEffect()
	if v == nil {
		t.Errorf("dml.NewCT_AlphaBiLevelEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AlphaBiLevelEffect should validate: %s", err)
	}
}

func TestCT_AlphaBiLevelEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AlphaBiLevelEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AlphaBiLevelEffect()
	xml.Unmarshal(buf, v2)
}
