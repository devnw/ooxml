package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_HeadersConstructor(t *testing.T) {
	v := dml.NewCT_Headers()
	if v == nil {
		t.Errorf("dml.NewCT_Headers must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Headers should validate: %s", err)
	}
}

func TestCT_HeadersMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Headers()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Headers()
	xml.Unmarshal(buf, v2)
}
