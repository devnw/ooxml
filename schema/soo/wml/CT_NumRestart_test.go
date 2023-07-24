package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_NumRestartConstructor(t *testing.T) {
	v := wml.NewCT_NumRestart()
	if v == nil {
		t.Errorf("wml.NewCT_NumRestart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_NumRestart should validate: %s", err)
	}
}

func TestCT_NumRestartMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_NumRestart()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_NumRestart()
	xml.Unmarshal(buf, v2)
}
