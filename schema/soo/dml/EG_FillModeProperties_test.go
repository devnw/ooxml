package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_FillModePropertiesConstructor(t *testing.T) {
	v := dml.NewEG_FillModeProperties()
	if v == nil {
		t.Errorf("dml.NewEG_FillModeProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_FillModeProperties should validate: %s", err)
	}
}

func TestEG_FillModePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_FillModeProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_FillModeProperties()
	xml.Unmarshal(buf, v2)
}
