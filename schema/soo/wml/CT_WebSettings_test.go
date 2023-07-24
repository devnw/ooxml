package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_WebSettingsConstructor(t *testing.T) {
	v := wml.NewCT_WebSettings()
	if v == nil {
		t.Errorf("wml.NewCT_WebSettings must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_WebSettings should validate: %s", err)
	}
}

func TestCT_WebSettingsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_WebSettings()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_WebSettings()
	xml.Unmarshal(buf, v2)
}
