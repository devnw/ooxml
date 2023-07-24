package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FtnEdnRefConstructor(t *testing.T) {
	v := wml.NewCT_FtnEdnRef()
	if v == nil {
		t.Errorf("wml.NewCT_FtnEdnRef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FtnEdnRef should validate: %s", err)
	}
}

func TestCT_FtnEdnRefMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FtnEdnRef()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FtnEdnRef()
	xml.Unmarshal(buf, v2)
}
