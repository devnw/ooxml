package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimateRotationBehaviorConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimateRotationBehavior()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimateRotationBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimateRotationBehavior should validate: %s", err)
	}
}

func TestCT_TLAnimateRotationBehaviorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimateRotationBehavior()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimateRotationBehavior()
	xml.Unmarshal(buf, v2)
}
