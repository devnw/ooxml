package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextBodyConstructor(t *testing.T) {
	v := dml.NewCT_TextBody()
	if v == nil {
		t.Errorf("dml.NewCT_TextBody must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextBody should validate: %s", err)
	}
}

func TestCT_TextBodyMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextBody()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextBody()
	xml.Unmarshal(buf, v2)
}
