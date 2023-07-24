package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_TextRunConstructor(t *testing.T) {
	v := dml.NewEG_TextRun()
	if v == nil {
		t.Errorf("dml.NewEG_TextRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextRun should validate: %s", err)
	}
}

func TestEG_TextRunMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextRun()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextRun()
	xml.Unmarshal(buf, v2)
}
