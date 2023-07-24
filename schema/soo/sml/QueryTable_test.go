package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestQueryTableConstructor(t *testing.T) {
	v := sml.NewQueryTable()
	if v == nil {
		t.Errorf("sml.NewQueryTable must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.QueryTable should validate: %s", err)
	}
}

func TestQueryTableMarshalUnmarshal(t *testing.T) {
	v := sml.NewQueryTable()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewQueryTable()
	xml.Unmarshal(buf, v2)
}
