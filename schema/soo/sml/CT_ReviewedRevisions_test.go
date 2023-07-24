package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ReviewedRevisionsConstructor(t *testing.T) {
	v := sml.NewCT_ReviewedRevisions()
	if v == nil {
		t.Errorf("sml.NewCT_ReviewedRevisions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ReviewedRevisions should validate: %s", err)
	}
}

func TestCT_ReviewedRevisionsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ReviewedRevisions()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ReviewedRevisions()
	xml.Unmarshal(buf, v2)
}
