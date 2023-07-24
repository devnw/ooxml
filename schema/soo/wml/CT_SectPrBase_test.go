package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SectPrBaseConstructor(t *testing.T) {
	v := wml.NewCT_SectPrBase()
	if v == nil {
		t.Errorf("wml.NewCT_SectPrBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SectPrBase should validate: %s", err)
	}
}

func TestCT_SectPrBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SectPrBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SectPrBase()
	xml.Unmarshal(buf, v2)
}
