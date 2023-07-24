package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_RPrBaseConstructor(t *testing.T) {
	v := wml.NewEG_RPrBase()
	if v == nil {
		t.Errorf("wml.NewEG_RPrBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RPrBase should validate: %s", err)
	}
}

func TestEG_RPrBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RPrBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RPrBase()
	xml.Unmarshal(buf, v2)
}
