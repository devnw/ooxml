package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextCharacterPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_TextCharacterProperties()
	if v == nil {
		t.Errorf("dml.NewCT_TextCharacterProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextCharacterProperties should validate: %s", err)
	}
}

func TestCT_TextCharacterPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextCharacterProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextCharacterProperties()
	xml.Unmarshal(buf, v2)
}
