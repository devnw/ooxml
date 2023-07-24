package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PermStartConstructor(t *testing.T) {
	v := wml.NewCT_PermStart()
	if v == nil {
		t.Errorf("wml.NewCT_PermStart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PermStart should validate: %s", err)
	}
}

func TestCT_PermStartMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PermStart()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PermStart()
	xml.Unmarshal(buf, v2)
}
