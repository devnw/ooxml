package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FilterColumnConstructor(t *testing.T) {
	v := sml.NewCT_FilterColumn()
	if v == nil {
		t.Errorf("sml.NewCT_FilterColumn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FilterColumn should validate: %s", err)
	}
}

func TestCT_FilterColumnMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FilterColumn()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FilterColumn()
	xml.Unmarshal(buf, v2)
}
