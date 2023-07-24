package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_QueryTableConstructor(t *testing.T) {
	v := sml.NewCT_QueryTable()
	if v == nil {
		t.Errorf("sml.NewCT_QueryTable must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_QueryTable should validate: %s", err)
	}
}

func TestCT_QueryTableMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_QueryTable()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_QueryTable()
	xml.Unmarshal(buf, v2)
}
