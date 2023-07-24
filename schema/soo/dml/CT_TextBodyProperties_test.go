package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextBodyPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_TextBodyProperties()
	if v == nil {
		t.Errorf("dml.NewCT_TextBodyProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextBodyProperties should validate: %s", err)
	}
}

func TestCT_TextBodyPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextBodyProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextBodyProperties()
	xml.Unmarshal(buf, v2)
}
