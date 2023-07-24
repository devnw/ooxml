package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_VolTopicConstructor(t *testing.T) {
	v := sml.NewCT_VolTopic()
	if v == nil {
		t.Errorf("sml.NewCT_VolTopic must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_VolTopic should validate: %s", err)
	}
}

func TestCT_VolTopicMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_VolTopic()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_VolTopic()
	xml.Unmarshal(buf, v2)
}
