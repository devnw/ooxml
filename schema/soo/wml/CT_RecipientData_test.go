package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RecipientDataConstructor(t *testing.T) {
	v := wml.NewCT_RecipientData()
	if v == nil {
		t.Errorf("wml.NewCT_RecipientData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RecipientData should validate: %s", err)
	}
}

func TestCT_RecipientDataMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RecipientData()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RecipientData()
	xml.Unmarshal(buf, v2)
}
