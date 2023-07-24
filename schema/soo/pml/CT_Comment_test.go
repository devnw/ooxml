package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CommentConstructor(t *testing.T) {
	v := pml.NewCT_Comment()
	if v == nil {
		t.Errorf("pml.NewCT_Comment must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Comment should validate: %s", err)
	}
}

func TestCT_CommentMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Comment()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Comment()
	xml.Unmarshal(buf, v2)
}
