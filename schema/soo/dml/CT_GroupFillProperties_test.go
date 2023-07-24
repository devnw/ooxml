package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GroupFillPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_GroupFillProperties()
	if v == nil {
		t.Errorf("dml.NewCT_GroupFillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GroupFillProperties should validate: %s", err)
	}
}

func TestCT_GroupFillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GroupFillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GroupFillProperties()
	xml.Unmarshal(buf, v2)
}
