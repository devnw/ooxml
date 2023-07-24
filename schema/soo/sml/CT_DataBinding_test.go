package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataBindingConstructor(t *testing.T) {
	v := sml.NewCT_DataBinding()
	if v == nil {
		t.Errorf("sml.NewCT_DataBinding must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataBinding should validate: %s", err)
	}
}

func TestCT_DataBindingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataBinding()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataBinding()
	xml.Unmarshal(buf, v2)
}
