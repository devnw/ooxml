package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LanguageConstructor(t *testing.T) {
	v := wml.NewCT_Language()
	if v == nil {
		t.Errorf("wml.NewCT_Language must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Language should validate: %s", err)
	}
}

func TestCT_LanguageMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Language()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Language()
	xml.Unmarshal(buf, v2)
}
