package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PPrConstructor(t *testing.T) {
	v := wml.NewCT_PPr()
	if v == nil {
		t.Errorf("wml.NewCT_PPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PPr should validate: %s", err)
	}
}

func TestCT_PPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PPr()
	xml.Unmarshal(buf, v2)
}
