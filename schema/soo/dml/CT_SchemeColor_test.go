package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_SchemeColorConstructor(t *testing.T) {
	v := dml.NewCT_SchemeColor()
	if v == nil {
		t.Errorf("dml.NewCT_SchemeColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_SchemeColor should validate: %s", err)
	}
}

func TestCT_SchemeColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_SchemeColor()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_SchemeColor()
	xml.Unmarshal(buf, v2)
}
