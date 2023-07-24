package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RecipientsConstructor(t *testing.T) {
	v := wml.NewCT_Recipients()
	if v == nil {
		t.Errorf("wml.NewCT_Recipients must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Recipients should validate: %s", err)
	}
}

func TestCT_RecipientsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Recipients()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Recipients()
	xml.Unmarshal(buf, v2)
}
