package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestEG_BackgroundConstructor(t *testing.T) {
	v := pml.NewEG_Background()
	if v == nil {
		t.Errorf("pml.NewEG_Background must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.EG_Background should validate: %s", err)
	}
}

func TestEG_BackgroundMarshalUnmarshal(t *testing.T) {
	v := pml.NewEG_Background()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewEG_Background()
	xml.Unmarshal(buf, v2)
}
