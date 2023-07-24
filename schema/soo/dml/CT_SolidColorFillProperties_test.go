package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_SolidColorFillPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_SolidColorFillProperties()
	if v == nil {
		t.Errorf("dml.NewCT_SolidColorFillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_SolidColorFillProperties should validate: %s", err)
	}
}

func TestCT_SolidColorFillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_SolidColorFillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_SolidColorFillProperties()
	xml.Unmarshal(buf, v2)
}
