package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWebSettingsConstructor(t *testing.T) {
	v := wml.NewWebSettings()
	if v == nil {
		t.Errorf("wml.NewWebSettings must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WebSettings should validate: %s", err)
	}
}

func TestWebSettingsMarshalUnmarshal(t *testing.T) {
	v := wml.NewWebSettings()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWebSettings()
	xml.Unmarshal(buf, v2)
}
