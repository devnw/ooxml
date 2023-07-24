package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_RPrConstructor(t *testing.T) {
	v := wml.NewEG_RPr()
	if v == nil {
		t.Errorf("wml.NewEG_RPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RPr should validate: %s", err)
	}
}

func TestEG_RPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RPr()
	xml.Unmarshal(buf, v2)
}
