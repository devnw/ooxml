package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellWatchesConstructor(t *testing.T) {
	v := sml.NewCT_CellWatches()
	if v == nil {
		t.Errorf("sml.NewCT_CellWatches must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellWatches should validate: %s", err)
	}
}

func TestCT_CellWatchesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellWatches()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellWatches()
	xml.Unmarshal(buf, v2)
}
