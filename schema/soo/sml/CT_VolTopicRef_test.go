package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_VolTopicRefConstructor(t *testing.T) {
	v := sml.NewCT_VolTopicRef()
	if v == nil {
		t.Errorf("sml.NewCT_VolTopicRef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_VolTopicRef should validate: %s", err)
	}
}

func TestCT_VolTopicRefMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_VolTopicRef()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_VolTopicRef()
	xml.Unmarshal(buf, v2)
}
