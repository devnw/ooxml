package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CustomerDataListConstructor(t *testing.T) {
	v := pml.NewCT_CustomerDataList()
	if v == nil {
		t.Errorf("pml.NewCT_CustomerDataList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CustomerDataList should validate: %s", err)
	}
}

func TestCT_CustomerDataListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CustomerDataList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CustomerDataList()
	xml.Unmarshal(buf, v2)
}
