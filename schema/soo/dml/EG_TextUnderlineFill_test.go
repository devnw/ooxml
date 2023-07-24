package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_TextUnderlineFillConstructor(t *testing.T) {
	v := dml.NewEG_TextUnderlineFill()
	if v == nil {
		t.Errorf("dml.NewEG_TextUnderlineFill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextUnderlineFill should validate: %s", err)
	}
}

func TestEG_TextUnderlineFillMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextUnderlineFill()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextUnderlineFill()
	xml.Unmarshal(buf, v2)
}
