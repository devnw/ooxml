package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SettingsConstructor(t *testing.T) {
	v := wml.NewCT_Settings()
	if v == nil {
		t.Errorf("wml.NewCT_Settings must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Settings should validate: %s", err)
	}
}

func TestCT_SettingsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Settings()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Settings()
	xml.Unmarshal(buf, v2)
}
