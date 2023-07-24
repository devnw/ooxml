package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AudioCDTimeConstructor(t *testing.T) {
	v := dml.NewCT_AudioCDTime()
	if v == nil {
		t.Errorf("dml.NewCT_AudioCDTime must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AudioCDTime should validate: %s", err)
	}
}

func TestCT_AudioCDTimeMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AudioCDTime()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AudioCDTime()
	xml.Unmarshal(buf, v2)
}
