package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FilterConstructor(t *testing.T) {
	v := sml.NewCT_Filter()
	if v == nil {
		t.Errorf("sml.NewCT_Filter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Filter should validate: %s", err)
	}
}

func TestCT_FilterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Filter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Filter()
	xml.Unmarshal(buf, v2)
}
