package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_NoFillPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_NoFillProperties()
	if v == nil {
		t.Errorf("dml.NewCT_NoFillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_NoFillProperties should validate: %s", err)
	}
}

func TestCT_NoFillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_NoFillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_NoFillProperties()
	xml.Unmarshal(buf, v2)
}
