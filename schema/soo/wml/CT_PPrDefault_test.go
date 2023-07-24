package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PPrDefaultConstructor(t *testing.T) {
	v := wml.NewCT_PPrDefault()
	if v == nil {
		t.Errorf("wml.NewCT_PPrDefault must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PPrDefault should validate: %s", err)
	}
}

func TestCT_PPrDefaultMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PPrDefault()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PPrDefault()
	xml.Unmarshal(buf, v2)
}
