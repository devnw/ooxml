package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimateColorBehaviorConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimateColorBehavior()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimateColorBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimateColorBehavior should validate: %s", err)
	}
}

func TestCT_TLAnimateColorBehaviorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimateColorBehavior()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimateColorBehavior()
	xml.Unmarshal(buf, v2)
}
