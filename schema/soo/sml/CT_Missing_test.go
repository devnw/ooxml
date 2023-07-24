package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MissingConstructor(t *testing.T) {
	v := sml.NewCT_Missing()
	if v == nil {
		t.Errorf("sml.NewCT_Missing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Missing should validate: %s", err)
	}
}

func TestCT_MissingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Missing()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Missing()
	xml.Unmarshal(buf, v2)
}
