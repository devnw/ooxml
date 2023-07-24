package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_PContentConstructor(t *testing.T) {
	v := wml.NewEG_PContent()
	if v == nil {
		t.Errorf("wml.NewEG_PContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_PContent should validate: %s", err)
	}
}

func TestEG_PContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_PContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_PContent()
	xml.Unmarshal(buf, v2)
}
