package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_FillPropertiesConstructor(t *testing.T) {
	v := dml.NewEG_FillProperties()
	if v == nil {
		t.Errorf("dml.NewEG_FillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_FillProperties should validate: %s", err)
	}
}

func TestEG_FillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_FillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_FillProperties()
	xml.Unmarshal(buf, v2)
}
