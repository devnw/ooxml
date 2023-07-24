package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_MathContentConstructor(t *testing.T) {
	v := wml.NewEG_MathContent()
	if v == nil {
		t.Errorf("wml.NewEG_MathContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_MathContent should validate: %s", err)
	}
}

func TestEG_MathContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_MathContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_MathContent()
	xml.Unmarshal(buf, v2)
}
