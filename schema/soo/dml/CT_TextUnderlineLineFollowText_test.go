package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextUnderlineLineFollowTextConstructor(t *testing.T) {
	v := dml.NewCT_TextUnderlineLineFollowText()
	if v == nil {
		t.Errorf("dml.NewCT_TextUnderlineLineFollowText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextUnderlineLineFollowText should validate: %s", err)
	}
}

func TestCT_TextUnderlineLineFollowTextMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextUnderlineLineFollowText()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextUnderlineLineFollowText()
	xml.Unmarshal(buf, v2)
}
