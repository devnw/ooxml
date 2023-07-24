package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GradientFillPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_GradientFillProperties()
	if v == nil {
		t.Errorf("dml.NewCT_GradientFillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GradientFillProperties should validate: %s", err)
	}
}

func TestCT_GradientFillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GradientFillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GradientFillProperties()
	xml.Unmarshal(buf, v2)
}
