package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SimpleFieldConstructor(t *testing.T) {
	v := wml.NewCT_SimpleField()
	if v == nil {
		t.Errorf("wml.NewCT_SimpleField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SimpleField should validate: %s", err)
	}
}

func TestCT_SimpleFieldMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SimpleField()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SimpleField()
	xml.Unmarshal(buf, v2)
}
