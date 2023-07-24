package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TrackChangeConstructor(t *testing.T) {
	v := wml.NewCT_TrackChange()
	if v == nil {
		t.Errorf("wml.NewCT_TrackChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TrackChange should validate: %s", err)
	}
}

func TestCT_TrackChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TrackChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TrackChange()
	xml.Unmarshal(buf, v2)
}
