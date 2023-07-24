package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTimeAnimateValueListConstructor(t *testing.T) {
	v := pml.NewCT_TLTimeAnimateValueList()
	if v == nil {
		t.Errorf("pml.NewCT_TLTimeAnimateValueList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTimeAnimateValueList should validate: %s", err)
	}
}

func TestCT_TLTimeAnimateValueListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTimeAnimateValueList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTimeAnimateValueList()
	xml.Unmarshal(buf, v2)
}
