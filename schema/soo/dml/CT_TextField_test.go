package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextFieldConstructor(t *testing.T) {
	v := dml.NewCT_TextField()
	if v == nil {
		t.Errorf("dml.NewCT_TextField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextField should validate: %s", err)
	}
}

func TestCT_TextFieldMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextField()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextField()
	xml.Unmarshal(buf, v2)
}
