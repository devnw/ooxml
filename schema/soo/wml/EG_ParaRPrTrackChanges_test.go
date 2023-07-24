package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_ParaRPrTrackChangesConstructor(t *testing.T) {
	v := wml.NewEG_ParaRPrTrackChanges()
	if v == nil {
		t.Errorf("wml.NewEG_ParaRPrTrackChanges must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_ParaRPrTrackChanges should validate: %s", err)
	}
}

func TestEG_ParaRPrTrackChangesMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_ParaRPrTrackChanges()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_ParaRPrTrackChanges()
	xml.Unmarshal(buf, v2)
}
