package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimateScaleBehaviorConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimateScaleBehavior()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimateScaleBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimateScaleBehavior should validate: %s", err)
	}
}

func TestCT_TLAnimateScaleBehaviorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimateScaleBehavior()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimateScaleBehavior()
	xml.Unmarshal(buf, v2)
}
