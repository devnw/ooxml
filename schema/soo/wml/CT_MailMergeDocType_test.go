package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MailMergeDocTypeConstructor(t *testing.T) {
	v := wml.NewCT_MailMergeDocType()
	if v == nil {
		t.Errorf("wml.NewCT_MailMergeDocType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MailMergeDocType should validate: %s", err)
	}
}

func TestCT_MailMergeDocTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MailMergeDocType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MailMergeDocType()
	xml.Unmarshal(buf, v2)
}
