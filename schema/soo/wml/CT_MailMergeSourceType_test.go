package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MailMergeSourceTypeConstructor(t *testing.T) {
	v := wml.NewCT_MailMergeSourceType()
	if v == nil {
		t.Errorf("wml.NewCT_MailMergeSourceType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MailMergeSourceType should validate: %s", err)
	}
}

func TestCT_MailMergeSourceTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MailMergeSourceType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MailMergeSourceType()
	xml.Unmarshal(buf, v2)
}
