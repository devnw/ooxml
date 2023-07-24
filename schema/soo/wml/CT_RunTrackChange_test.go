package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RunTrackChangeConstructor(t *testing.T) {
	v := wml.NewCT_RunTrackChange()
	if v == nil {
		t.Errorf("wml.NewCT_RunTrackChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RunTrackChange should validate: %s", err)
	}
}

func TestCT_RunTrackChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RunTrackChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RunTrackChange()
	xml.Unmarshal(buf, v2)
}
