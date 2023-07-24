package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CellMergeTrackChangeConstructor(t *testing.T) {
	v := wml.NewCT_CellMergeTrackChange()
	if v == nil {
		t.Errorf("wml.NewCT_CellMergeTrackChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CellMergeTrackChange should validate: %s", err)
	}
}

func TestCT_CellMergeTrackChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CellMergeTrackChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CellMergeTrackChange()
	xml.Unmarshal(buf, v2)
}
