package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_FillPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_FillProperties()
	if v == nil {
		t.Errorf("dml.NewCT_FillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FillProperties should validate: %s", err)
	}
}

func TestCT_FillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FillProperties()
	xml.Unmarshal(buf, v2)
}
