package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TcBordersConstructor(t *testing.T) {
	v := wml.NewCT_TcBorders()
	if v == nil {
		t.Errorf("wml.NewCT_TcBorders must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TcBorders should validate: %s", err)
	}
}

func TestCT_TcBordersMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TcBorders()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TcBorders()
	xml.Unmarshal(buf, v2)
}
