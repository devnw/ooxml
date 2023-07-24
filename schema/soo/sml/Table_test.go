package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestTableConstructor(t *testing.T) {
	v := sml.NewTable()
	if v == nil {
		t.Errorf("sml.NewTable must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Table should validate: %s", err)
	}
}

func TestTableMarshalUnmarshal(t *testing.T) {
	v := sml.NewTable()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewTable()
	xml.Unmarshal(buf, v2)
}
