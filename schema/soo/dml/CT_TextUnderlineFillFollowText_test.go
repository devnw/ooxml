package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextUnderlineFillFollowTextConstructor(t *testing.T) {
	v := dml.NewCT_TextUnderlineFillFollowText()
	if v == nil {
		t.Errorf("dml.NewCT_TextUnderlineFillFollowText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextUnderlineFillFollowText should validate: %s", err)
	}
}

func TestCT_TextUnderlineFillFollowTextMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextUnderlineFillFollowText()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextUnderlineFillFollowText()
	xml.Unmarshal(buf, v2)
}
