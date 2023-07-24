package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TableColumnConstructor(t *testing.T) {
	v := sml.NewCT_TableColumn()
	if v == nil {
		t.Errorf("sml.NewCT_TableColumn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TableColumn should validate: %s", err)
	}
}

func TestCT_TableColumnMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TableColumn()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TableColumn()
	xml.Unmarshal(buf, v2)
}
