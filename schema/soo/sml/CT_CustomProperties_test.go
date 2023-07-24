package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomPropertiesConstructor(t *testing.T) {
	v := sml.NewCT_CustomProperties()
	if v == nil {
		t.Errorf("sml.NewCT_CustomProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomProperties should validate: %s", err)
	}
}

func TestCT_CustomPropertiesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomProperties()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomProperties()
	xml.Unmarshal(buf, v2)
}
