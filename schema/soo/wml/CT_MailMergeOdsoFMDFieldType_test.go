package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MailMergeOdsoFMDFieldTypeConstructor(t *testing.T) {
	v := wml.NewCT_MailMergeOdsoFMDFieldType()
	if v == nil {
		t.Errorf("wml.NewCT_MailMergeOdsoFMDFieldType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MailMergeOdsoFMDFieldType should validate: %s", err)
	}
}

func TestCT_MailMergeOdsoFMDFieldTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MailMergeOdsoFMDFieldType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MailMergeOdsoFMDFieldType()
	xml.Unmarshal(buf, v2)
}
