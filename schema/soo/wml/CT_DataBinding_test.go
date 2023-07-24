package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DataBindingConstructor(t *testing.T) {
	v := wml.NewCT_DataBinding()
	if v == nil {
		t.Errorf("wml.NewCT_DataBinding must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DataBinding should validate: %s", err)
	}
}

func TestCT_DataBindingMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DataBinding()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DataBinding()
	xml.Unmarshal(buf, v2)
}
