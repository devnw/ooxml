package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartBehaviorConstructor(t *testing.T) {
	v := wml.NewCT_DocPartBehavior()
	if v == nil {
		t.Errorf("wml.NewCT_DocPartBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPartBehavior should validate: %s", err)
	}
}

func TestCT_DocPartBehaviorMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPartBehavior()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPartBehavior()
	xml.Unmarshal(buf, v2)
}
