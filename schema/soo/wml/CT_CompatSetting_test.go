package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CompatSettingConstructor(t *testing.T) {
	v := wml.NewCT_CompatSetting()
	if v == nil {
		t.Errorf("wml.NewCT_CompatSetting must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CompatSetting should validate: %s", err)
	}
}

func TestCT_CompatSettingMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CompatSetting()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CompatSetting()
	xml.Unmarshal(buf, v2)
}
