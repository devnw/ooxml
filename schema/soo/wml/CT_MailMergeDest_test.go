package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MailMergeDestConstructor(t *testing.T) {
	v := wml.NewCT_MailMergeDest()
	if v == nil {
		t.Errorf("wml.NewCT_MailMergeDest must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MailMergeDest should validate: %s", err)
	}
}

func TestCT_MailMergeDestMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MailMergeDest()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MailMergeDest()
	xml.Unmarshal(buf, v2)
}
