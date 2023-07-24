package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_UndoInfoConstructor(t *testing.T) {
	v := sml.NewCT_UndoInfo()
	if v == nil {
		t.Errorf("sml.NewCT_UndoInfo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_UndoInfo should validate: %s", err)
	}
}

func TestCT_UndoInfoMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_UndoInfo()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_UndoInfo()
	xml.Unmarshal(buf, v2)
}
