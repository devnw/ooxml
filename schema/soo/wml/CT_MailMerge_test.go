package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MailMergeConstructor(t *testing.T) {
	v := wml.NewCT_MailMerge()
	if v == nil {
		t.Errorf("wml.NewCT_MailMerge must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MailMerge should validate: %s", err)
	}
}

func TestCT_MailMergeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MailMerge()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MailMerge()
	xml.Unmarshal(buf, v2)
}
