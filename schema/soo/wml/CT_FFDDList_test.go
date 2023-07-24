package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFDDListConstructor(t *testing.T) {
	v := wml.NewCT_FFDDList()
	if v == nil {
		t.Errorf("wml.NewCT_FFDDList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFDDList should validate: %s", err)
	}
}

func TestCT_FFDDListMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFDDList()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFDDList()
	xml.Unmarshal(buf, v2)
}
