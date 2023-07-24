package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FileSharingConstructor(t *testing.T) {
	v := sml.NewCT_FileSharing()
	if v == nil {
		t.Errorf("sml.NewCT_FileSharing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FileSharing should validate: %s", err)
	}
}

func TestCT_FileSharingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FileSharing()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FileSharing()
	xml.Unmarshal(buf, v2)
}
