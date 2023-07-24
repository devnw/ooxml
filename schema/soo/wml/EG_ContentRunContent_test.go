package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_ContentRunContentConstructor(t *testing.T) {
	v := wml.NewEG_ContentRunContent()
	if v == nil {
		t.Errorf("wml.NewEG_ContentRunContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_ContentRunContent should validate: %s", err)
	}
}

func TestEG_ContentRunContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_ContentRunContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_ContentRunContent()
	xml.Unmarshal(buf, v2)
}
