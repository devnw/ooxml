package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestNumberingConstructor(t *testing.T) {
	v := wml.NewNumbering()
	if v == nil {
		t.Errorf("wml.NewNumbering must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Numbering should validate: %s", err)
	}
}

func TestNumberingMarshalUnmarshal(t *testing.T) {
	v := wml.NewNumbering()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewNumbering()
	xml.Unmarshal(buf, v2)
}
