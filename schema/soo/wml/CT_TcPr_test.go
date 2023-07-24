package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TcPrConstructor(t *testing.T) {
	v := wml.NewCT_TcPr()
	if v == nil {
		t.Errorf("wml.NewCT_TcPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TcPr should validate: %s", err)
	}
}

func TestCT_TcPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TcPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TcPr()
	xml.Unmarshal(buf, v2)
}
