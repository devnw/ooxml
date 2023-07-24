package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ReadingModeInkLockDownConstructor(t *testing.T) {
	v := wml.NewCT_ReadingModeInkLockDown()
	if v == nil {
		t.Errorf("wml.NewCT_ReadingModeInkLockDown must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ReadingModeInkLockDown should validate: %s", err)
	}
}

func TestCT_ReadingModeInkLockDownMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ReadingModeInkLockDown()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ReadingModeInkLockDown()
	xml.Unmarshal(buf, v2)
}
