package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TcPrInnerConstructor(t *testing.T) {
	v := wml.NewCT_TcPrInner()
	if v == nil {
		t.Errorf("wml.NewCT_TcPrInner must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TcPrInner should validate: %s", err)
	}
}

func TestCT_TcPrInnerMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TcPrInner()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TcPrInner()
	xml.Unmarshal(buf, v2)
}
