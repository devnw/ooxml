package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestSettingsConstructor(t *testing.T) {
	v := wml.NewSettings()
	if v == nil {
		t.Errorf("wml.NewSettings must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Settings should validate: %s", err)
	}
}

func TestSettingsMarshalUnmarshal(t *testing.T) {
	v := wml.NewSettings()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewSettings()
	xml.Unmarshal(buf, v2)
}
