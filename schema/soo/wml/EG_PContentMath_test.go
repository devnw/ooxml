package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_PContentMathConstructor(t *testing.T) {
	v := wml.NewEG_PContentMath()
	if v == nil {
		t.Errorf("wml.NewEG_PContentMath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_PContentMath should validate: %s", err)
	}
}

func TestEG_PContentMathMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_PContentMath()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_PContentMath()
	xml.Unmarshal(buf, v2)
}
