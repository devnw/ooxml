package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestAG_AutoFormatConstructor(t *testing.T) {
	v := sml.NewAG_AutoFormat()
	if v == nil {
		t.Errorf("sml.NewAG_AutoFormat must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.AG_AutoFormat should validate: %s", err)
	}
}

func TestAG_AutoFormatMarshalUnmarshal(t *testing.T) {
	v := sml.NewAG_AutoFormat()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewAG_AutoFormat()
	xml.Unmarshal(buf, v2)
}
