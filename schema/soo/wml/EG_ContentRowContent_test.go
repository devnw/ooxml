package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_ContentRowContentConstructor(t *testing.T) {
	v := wml.NewEG_ContentRowContent()
	if v == nil {
		t.Errorf("wml.NewEG_ContentRowContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_ContentRowContent should validate: %s", err)
	}
}

func TestEG_ContentRowContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_ContentRowContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_ContentRowContent()
	xml.Unmarshal(buf, v2)
}
