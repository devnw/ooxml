package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocPartBehaviorsConstructor(t *testing.T) {
	v := wml.NewCT_DocPartBehaviors()
	if v == nil {
		t.Errorf("wml.NewCT_DocPartBehaviors must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocPartBehaviors should validate: %s", err)
	}
}

func TestCT_DocPartBehaviorsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocPartBehaviors()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocPartBehaviors()
	xml.Unmarshal(buf, v2)
}
