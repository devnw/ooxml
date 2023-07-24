package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CalendarTypeConstructor(t *testing.T) {
	v := wml.NewCT_CalendarType()
	if v == nil {
		t.Errorf("wml.NewCT_CalendarType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CalendarType should validate: %s", err)
	}
}

func TestCT_CalendarTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CalendarType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CalendarType()
	xml.Unmarshal(buf, v2)
}
