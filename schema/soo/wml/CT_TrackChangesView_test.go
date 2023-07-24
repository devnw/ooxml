package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TrackChangesViewConstructor(t *testing.T) {
	v := wml.NewCT_TrackChangesView()
	if v == nil {
		t.Errorf("wml.NewCT_TrackChangesView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TrackChangesView should validate: %s", err)
	}
}

func TestCT_TrackChangesViewMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TrackChangesView()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TrackChangesView()
	xml.Unmarshal(buf, v2)
}
