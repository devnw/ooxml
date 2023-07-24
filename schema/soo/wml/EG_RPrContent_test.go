package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_RPrContentConstructor(t *testing.T) {
	v := wml.NewEG_RPrContent()
	if v == nil {
		t.Errorf("wml.NewEG_RPrContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RPrContent should validate: %s", err)
	}
}

func TestEG_RPrContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RPrContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RPrContent()
	xml.Unmarshal(buf, v2)
}
