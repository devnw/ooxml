package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FtnEdnConstructor(t *testing.T) {
	v := wml.NewCT_FtnEdn()
	if v == nil {
		t.Errorf("wml.NewCT_FtnEdn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FtnEdn should validate: %s", err)
	}
}

func TestCT_FtnEdnMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FtnEdn()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FtnEdn()
	xml.Unmarshal(buf, v2)
}
