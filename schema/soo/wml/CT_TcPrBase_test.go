package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TcPrBaseConstructor(t *testing.T) {
	v := wml.NewCT_TcPrBase()
	if v == nil {
		t.Errorf("wml.NewCT_TcPrBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TcPrBase should validate: %s", err)
	}
}

func TestCT_TcPrBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TcPrBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TcPrBase()
	xml.Unmarshal(buf, v2)
}
