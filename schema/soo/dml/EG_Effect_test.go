package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_EffectConstructor(t *testing.T) {
	v := dml.NewEG_Effect()
	if v == nil {
		t.Errorf("dml.NewEG_Effect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_Effect should validate: %s", err)
	}
}

func TestEG_EffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_Effect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_Effect()
	xml.Unmarshal(buf, v2)
}
