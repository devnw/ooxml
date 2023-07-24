package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TablePropertiesChoiceConstructor(t *testing.T) {
	v := dml.NewCT_TablePropertiesChoice()
	if v == nil {
		t.Errorf("dml.NewCT_TablePropertiesChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TablePropertiesChoice should validate: %s", err)
	}
}

func TestCT_TablePropertiesChoiceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TablePropertiesChoice()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TablePropertiesChoice()
	xml.Unmarshal(buf, v2)
}
