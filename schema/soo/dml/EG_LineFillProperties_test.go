package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_LineFillPropertiesConstructor(t *testing.T) {
	v := dml.NewEG_LineFillProperties()
	if v == nil {
		t.Errorf("dml.NewEG_LineFillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_LineFillProperties should validate: %s", err)
	}
}

func TestEG_LineFillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_LineFillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_LineFillProperties()
	xml.Unmarshal(buf, v2)
}
