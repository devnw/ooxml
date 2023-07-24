package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestRecipientsConstructor(t *testing.T) {
	v := wml.NewRecipients()
	if v == nil {
		t.Errorf("wml.NewRecipients must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Recipients should validate: %s", err)
	}
}

func TestRecipientsMarshalUnmarshal(t *testing.T) {
	v := wml.NewRecipients()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewRecipients()
	xml.Unmarshal(buf, v2)
}
