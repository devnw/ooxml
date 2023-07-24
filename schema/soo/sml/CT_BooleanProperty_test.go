package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_BooleanPropertyConstructor(t *testing.T) {
	v := sml.NewCT_BooleanProperty()
	if v == nil {
		t.Errorf("sml.NewCT_BooleanProperty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_BooleanProperty should validate: %s", err)
	}
}

func TestCT_BooleanPropertyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_BooleanProperty()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_BooleanProperty()
	xml.Unmarshal(buf, v2)
}
