package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_ColorChoiceConstructor(t *testing.T) {
	v := dml.NewEG_ColorChoice()
	if v == nil {
		t.Errorf("dml.NewEG_ColorChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_ColorChoice should validate: %s", err)
	}
}

func TestEG_ColorChoiceMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_ColorChoice()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_ColorChoice()
	xml.Unmarshal(buf, v2)
}
