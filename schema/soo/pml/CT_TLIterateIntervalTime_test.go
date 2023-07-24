package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLIterateIntervalTimeConstructor(t *testing.T) {
	v := pml.NewCT_TLIterateIntervalTime()
	if v == nil {
		t.Errorf("pml.NewCT_TLIterateIntervalTime must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLIterateIntervalTime should validate: %s", err)
	}
}

func TestCT_TLIterateIntervalTimeMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLIterateIntervalTime()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLIterateIntervalTime()
	xml.Unmarshal(buf, v2)
}
