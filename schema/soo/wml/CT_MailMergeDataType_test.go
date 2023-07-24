package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MailMergeDataTypeConstructor(t *testing.T) {
	v := wml.NewCT_MailMergeDataType()
	if v == nil {
		t.Errorf("wml.NewCT_MailMergeDataType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MailMergeDataType should validate: %s", err)
	}
}

func TestCT_MailMergeDataTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MailMergeDataType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MailMergeDataType()
	xml.Unmarshal(buf, v2)
}
