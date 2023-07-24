package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestFtrConstructor(t *testing.T) {
	v := wml.NewFtr()
	if v == nil {
		t.Errorf("wml.NewFtr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Ftr should validate: %s", err)
	}
}

func TestFtrMarshalUnmarshal(t *testing.T) {
	v := wml.NewFtr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewFtr()
	xml.Unmarshal(buf, v2)
}
