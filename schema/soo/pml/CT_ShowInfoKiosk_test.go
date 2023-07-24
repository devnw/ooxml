package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ShowInfoKioskConstructor(t *testing.T) {
	v := pml.NewCT_ShowInfoKiosk()
	if v == nil {
		t.Errorf("pml.NewCT_ShowInfoKiosk must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ShowInfoKiosk should validate: %s", err)
	}
}

func TestCT_ShowInfoKioskMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ShowInfoKiosk()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ShowInfoKiosk()
	xml.Unmarshal(buf, v2)
}
