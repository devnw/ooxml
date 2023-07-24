package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TcMarConstructor(t *testing.T) {
	v := wml.NewCT_TcMar()
	if v == nil {
		t.Errorf("wml.NewCT_TcMar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TcMar should validate: %s", err)
	}
}

func TestCT_TcMarMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TcMar()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TcMar()
	xml.Unmarshal(buf, v2)
}
