package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_ContentBlockContentConstructor(t *testing.T) {
	v := wml.NewEG_ContentBlockContent()
	if v == nil {
		t.Errorf("wml.NewEG_ContentBlockContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_ContentBlockContent should validate: %s", err)
	}
}

func TestEG_ContentBlockContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_ContentBlockContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_ContentBlockContent()
	xml.Unmarshal(buf, v2)
}
