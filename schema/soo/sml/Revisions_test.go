package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestRevisionsConstructor(t *testing.T) {
	v := sml.NewRevisions()
	if v == nil {
		t.Errorf("sml.NewRevisions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Revisions should validate: %s", err)
	}
}

func TestRevisionsMarshalUnmarshal(t *testing.T) {
	v := sml.NewRevisions()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewRevisions()
	xml.Unmarshal(buf, v2)
}
