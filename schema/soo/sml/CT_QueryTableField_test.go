package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_QueryTableFieldConstructor(t *testing.T) {
	v := sml.NewCT_QueryTableField()
	if v == nil {
		t.Errorf("sml.NewCT_QueryTableField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_QueryTableField should validate: %s", err)
	}
}

func TestCT_QueryTableFieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_QueryTableField()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_QueryTableField()
	xml.Unmarshal(buf, v2)
}
