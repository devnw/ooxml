package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PPrBaseConstructor(t *testing.T) {
	v := wml.NewCT_PPrBase()
	if v == nil {
		t.Errorf("wml.NewCT_PPrBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PPrBase should validate: %s", err)
	}
}

func TestCT_PPrBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PPrBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PPrBase()
	xml.Unmarshal(buf, v2)
}
