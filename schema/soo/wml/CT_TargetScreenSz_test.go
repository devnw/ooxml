package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TargetScreenSzConstructor(t *testing.T) {
	v := wml.NewCT_TargetScreenSz()
	if v == nil {
		t.Errorf("wml.NewCT_TargetScreenSz must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TargetScreenSz should validate: %s", err)
	}
}

func TestCT_TargetScreenSzMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TargetScreenSz()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TargetScreenSz()
	xml.Unmarshal(buf, v2)
}
