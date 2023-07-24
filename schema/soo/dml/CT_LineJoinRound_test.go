package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_LineJoinRoundConstructor(t *testing.T) {
	v := dml.NewCT_LineJoinRound()
	if v == nil {
		t.Errorf("dml.NewCT_LineJoinRound must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_LineJoinRound should validate: %s", err)
	}
}

func TestCT_LineJoinRoundMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_LineJoinRound()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_LineJoinRound()
	xml.Unmarshal(buf, v2)
}
