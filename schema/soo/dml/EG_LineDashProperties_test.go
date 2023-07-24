package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_LineDashPropertiesConstructor(t *testing.T) {
	v := dml.NewEG_LineDashProperties()
	if v == nil {
		t.Errorf("dml.NewEG_LineDashProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_LineDashProperties should validate: %s", err)
	}
}

func TestEG_LineDashPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_LineDashProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_LineDashProperties()
	xml.Unmarshal(buf, v2)
}
