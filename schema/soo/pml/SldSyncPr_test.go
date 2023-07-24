package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestSldSyncPrConstructor(t *testing.T) {
	v := pml.NewSldSyncPr()
	if v == nil {
		t.Errorf("pml.NewSldSyncPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.SldSyncPr should validate: %s", err)
	}
}

func TestSldSyncPrMarshalUnmarshal(t *testing.T) {
	v := pml.NewSldSyncPr()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewSldSyncPr()
	xml.Unmarshal(buf, v2)
}
